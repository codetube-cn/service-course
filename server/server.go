package server

import "codetube.cn/proto/service_course"

type CourseRegisterServer struct {
	service_course.UnimplementedCourseServer
}

func NewCourseRegisterServer() *CourseRegisterServer {
	return &CourseRegisterServer{}
}
