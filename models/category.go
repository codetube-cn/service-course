package models

import (
	"gorm.io/gorm"
)

// Category 课程分类模型
type Category struct {
	gorm.Model
	ParentId    uint   //上级分类
	Name        string //分类名称
	Ename       string //分类英文名称
	Title       string //分类页面标题
	Keywords    string //分类页面 keywords
	Description string //分类页面 description
	Enabled     int    //是否启用
	Sort        int    //排序序号
}

// Categories 课程分类列表
type Categories []*Category
