package server

import (
	"codetube.cn/proto/service_course"
	"codetube.cn/service-source/logic"
	"context"
)

// GetCategoryTree 获取课程分类树
func (s *CourseRegisterServer) GetCategoryTree(c context.Context, request *service_course.CategoriesTreeRequest) (*service_course.CategoriesTree, error) {
	categories := logic.GetAllCategoriesSorted()
	categoryTree := logic.GetCategoryTree(0, categories)
	return &service_course.CategoriesTree{Data: categoryTree}, nil
}
