package models

import (
	"gorm.io/gorm"
	"time"
)

// Course 课程模型
type Course struct {
	gorm.Model
	UserId         string
	Name           string
	CategoryId     uint
	UrlName        string
	Cover          string
	Introduction   string
	Price          uint
	ViewsTotal     uint
	ChaptersTotal  uint
	VideoTimeTotal uint
	PublishedAt    time.Time
	RecommendedAt  time.Time
}
