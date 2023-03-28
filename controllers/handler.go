package controllers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/newhorizon-tech-vn/redis-example/controllers/v1"
)

type IHandler interface {
	GetAlbums(c *gin.Context)
	PostAlbums(c *gin.Context)
}

func MakeHandler(version string) IHandler {
	return &v1.Handler{}
}
