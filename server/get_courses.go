package server

import (
	"codetube.cn/core/codes"
	"codetube.cn/core/errors"
	"codetube.cn/core/libraries"
	"codetube.cn/core/vars"
	"codetube.cn/proto/service_course"
	"codetube.cn/service-source/components"
	"codetube.cn/service-source/models"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strconv"
)

// GetCourses 获取课程分类树
func (s *CourseRegisterServer) GetCourses(c context.Context, request *service_course.GetCoursesRequest) (*service_course.CourseList, error) {
	courses := make([]*models.Course, 0)
	page := request.GetPage()
	if page < 1 {
		page = 1
	}
	pageSize := request.GetPageSize()
	if pageSize < 1 {
		pageSize = vars.PageSize
	}
	query := components.CourseDB.Limit(int(pageSize)).Offset(int((page - 1) * pageSize))
	//分类
	if request.GetCategoryId() > 0 {
		query.Where("category_id = ?", request.GetCategoryId())
	}
	//推荐课程
	if request.GetIsRecommended() {
		query.Order("recommended_at desc")
		//如果不自动填充，直接 where 条件限定只能是设置为推荐的课程
		if !request.GetIsRecommendedFill() {
			query.Where("recommended_at is not null")
		}
		query.Order("published_at desc").Order("created_at desc")
	}
	query.Where("published_at is not null")
	result := query.Find(&courses)
	if result.Error != nil {
		log.Println("[err:"+strconv.Itoa(codes.SqlFail)+"]查询课程列表：", result.Error)
		return nil, errors.Wrap("查询课程列表出错", result.Error)
	}

	courseList := &service_course.CourseList{Courses: make([]*service_course.CourseModel, 0)}

	if len(courses) > 0 {
		for _, course := range courses {
			courseList.Courses = append(courseList.Courses, &service_course.CourseModel{
				Id:             int64(course.ID),
				UserId:         course.UserId,
				Name:           course.Name,
				CategoryId:     int64(course.CategoryId),
				UrlName:        course.UrlName,
				Cover:          course.Cover,
				Introduction:   course.Introduction,
				Price:          float32(libraries.PriceCentToYuan(int64(course.Price))),
				ViewsTotal:     int64(course.ViewsTotal),
				ChaptersTotal:  int64(course.ChaptersTotal),
				VideoTimeTotal: int64(course.VideoTimeTotal),
				RecommendedAt:  timestamppb.New(course.RecommendedAt),
				PublishedAt:    timestamppb.New(course.PublishedAt),
			})
		}
	}

	return courseList, nil
}
