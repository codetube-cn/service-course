package server

import (
	"codetube.cn/proto/service_course"
	"codetube.cn/service-source/logic"
	"codetube.cn/service-source/models"
	"context"
)

// GetCategoryById 根据 ID 获取课程分类信息
func (s *CourseRegisterServer) GetCategoryById(c context.Context, request *service_course.GetCategoryByIdRequest) (*service_course.Category, error) {
	return logic.GetCategory(&models.Category{
		ID: uint(request.GetId()),
	})
}
