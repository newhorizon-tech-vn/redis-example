package models

import "github.com/newhorizon-tech-vn/redis-example/models/entities"

func GetClassessByTeacherId(userId int) ([]*entities.ClassTeacher, error) {
	var result []*entities.ClassTeacher
	err := DBConnection.Preload("Class").Where("user_UserId = ?", userId).Find(&result).Error
	return result, err
}

func GetClassessByStudentId(userId int) ([]*entities.ClassStudent, error) {
	var result []*entities.ClassStudent
	err := DBConnection.Preload("Class").Where("user_UserId = ?", userId).Find(&result).Error
	return result, err
}
