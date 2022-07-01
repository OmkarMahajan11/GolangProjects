package controller

import (
	"github.com/gin-gonic/gin"
	"poc.com/gin-poc/entity"
	"poc.com/gin-poc/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type videoController struct {
	service service.VideoService
}

func New() VideoController {
	return videoController{
		service: service.New(),
	}
}

func (vc videoController) FindAll() []entity.Video {
	return vc.service.FindAll()
}

func (vc videoController) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	return vc.service.Save(video)
}
