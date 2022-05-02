package server

import (
	"codetube.cn/proto/service_course"
	"codetube.cn/service-source/logic"
	"codetube.cn/service-source/models"
	"context"
)

// GetCategoryByUrlName 根据 UrlName 获取课程分类信息
func (s *CourseRegisterServer) GetCategoryByUrlName(c context.Context, request *service_course.GetCategoryByUrlNameRequest) (*service_course.Category, error) {
	return logic.GetCategory(&models.Category{
		UrlName: request.GetUrlName(),
	})
}
