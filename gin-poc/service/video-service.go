package service

import "poc.com/gin-poc/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(v entity.Video) entity.Video {
	service.videos = append(service.videos, v)
	return v
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
