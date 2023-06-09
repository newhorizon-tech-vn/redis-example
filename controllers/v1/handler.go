package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/newhorizon-tech-vn/redis-example/middleware/authorize"
	"github.com/newhorizon-tech-vn/redis-example/services"
	"golang.org/x/exp/slices"
)

type Handler struct{}

func (h *Handler) CheckClass(c *gin.Context) {
	classId, err := strconv.Atoi(c.Param("classId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	userId, exists := c.Get(authorize.KeyUserId)
	if exists == false {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	roleId, exists := c.Get(authorize.KeyRoleId)
	if exists == false {
		c.JSON(http.StatusNonAuthoritativeInfo, nil)
		return
	}

	classIds, err := (&services.ClassService{}).GetClassIds(userId.(int), roleId.(int))
	index := slices.IndexFunc(classIds, func(id int) bool { return id == classId })
	if index < 0 {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}
