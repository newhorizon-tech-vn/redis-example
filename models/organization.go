package models

import (
	"github.com/newhorizon-tech-vn/redis-example/models/entities"
)

func GetOrganizationWithClasses(organizationId int) (*entities.Organization, error) {
	result := &entities.Organization{}
	err := DBConnection.Preload("Schools.Classes").First(&result, organizationId).Error
	return result, err
}
