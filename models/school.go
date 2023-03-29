package models

import (
	"time"

	"github.com/newhorizon-tech-vn/redis-example/models/entities"
)

func GetUserOrgSchoolsWithClassesByUserId(userId int) ([]*entities.UserOrgSchool, error) {
	var result []*entities.UserOrgSchool
	err := DBConnection.Preload("School.Classes").Where("UserId = ?", userId).Find(&result).Error
	return result, err
}

func GetOrganizationByUserId(userId int) (*entities.Organization, error) {
	result := &entities.UserOrgSchool{}
	err := DBConnection.Preload("Organization").Where("UserId = ? AND EndDate > ?", userId, time.Now().UTC()).First(&result).Error
	return result.GetOrganization(), err
}
