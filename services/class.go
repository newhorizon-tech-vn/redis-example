package services

import (
	"context"
	"fmt"
	"time"

	"github.com/newhorizon-tech-vn/redis-example/cache"
	"github.com/newhorizon-tech-vn/redis-example/models"
	"github.com/newhorizon-tech-vn/redis-example/models/entities"
	"github.com/newhorizon-tech-vn/redis-example/pkg/util"
	"golang.org/x/exp/maps"
)

type ClassService struct {
	ClassId int
}

func (*ClassService) GetClassIds(userId, userRoleId int) ([]int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*60))
	defer cancel()

	// try to get from cache
	result, err := cache.GetClassIdsOfUser(ctx, userId)
	if err != nil {
		return result, nil
	}

	// try to get data from storage
	if userRoleId == entities.ROLE_STUDENT {
		rs, err := models.GetClassessByStudentId(userId)
		if err != nil {
			return nil, err
		}

		result = util.MapFunc(rs, func(c *entities.ClassStudent) int {
			return c.ClassId
		})
	} else if userRoleId == entities.ROLE_TEACHER {
		rs, err := models.GetClassessByTeacherId(userId)
		if err != nil {
			return nil, err
		}

		result = util.MapFunc(rs, func(c *entities.ClassTeacher) int {
			return c.ClassId
		})

	} else if userRoleId == entities.ROLE_ORGANIZATION_ADMIN {
		organization, err := models.GetOrganizationByUserId(userId)
		if err != nil {
			return nil, err
		}

		rs, err := models.GetOrganizationWithClasses(organization.GetId())
		if err != nil {
			return nil, err
		}

		classIds := make(map[int]bool)
		for _, school := range rs.GetSchools() {
			for _, class := range school.GetClasses() {
				classIds[class.ClassId] = true
			}
		}

		result = maps.Keys(classIds)

	} else if userRoleId == entities.ROLE_SCHOOL_ADMIN {
		rs, err := models.GetUserOrgSchoolsWithClassesByUserId(userId)
		if err != nil {
			return nil, err
		}

		classIds := make(map[int]bool)
		for _, val := range rs {
			for _, class := range val.GetSchool().GetClasses() {
				classIds[class.ClassId] = true
			}
		}

		result = maps.Keys(classIds)

	} else {
		return nil, fmt.Errorf("role invalid")
	}

	// update cache
	cache.SetClassIdsOfUser(ctx, userId, result)

	// return result
	return result, nil
}
