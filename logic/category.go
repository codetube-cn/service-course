package logic

import (
	"codetube.cn/core/codes"
	"codetube.cn/proto/service_course"
	"codetube.cn/service-source/components"
	"codetube.cn/service-source/models"
	"log"
	"strconv"
)

// GetAllCategoriesSorted 获取所有正常使用的部门，并按父部门及 sort 排序
func GetAllCategoriesSorted() []*models.Category {
	categories := make([]*models.Category, 0)
	result := components.CourseDB.Where("enabled = ?", 1).Order("parent_id asc").Order("sort asc").Find(&categories)
	if result.Error != nil {
		log.Println("[err:"+strconv.Itoa(codes.SqlFail)+"]查询课程分类列表：", result.Error)
	}
	return categories
}

// GetCategoryTree 获取分类树结构
func GetCategoryTree(parentId uint, allCategories []*models.Category) []*service_course.CategoryTree {
	categoryTree := make([]*service_course.CategoryTree, 0)
	for _, v := range allCategories {
		if v.ParentId == parentId {
			categoryTree = append(categoryTree, &service_course.CategoryTree{
				Id:       int64(v.ID),
				ParentId: int64(v.ParentId),
				Name:     v.Name,
				Ename:    v.Ename,
				Title:    v.Title,
				Children: GetCategoryTree(v.ID, allCategories),
			})
		}
	}
	return categoryTree
}
