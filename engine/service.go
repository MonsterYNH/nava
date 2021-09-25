package engine

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Name() string
	Middlewares() []gin.HandlerFunc
	RegisterHandler(*Engine, *gin.RouterGroup) error
}

type UnimplementedService struct{}

func (service *UnimplementedService) Name() string {
	return "UnimplementedServiceName"
}

func (service *UnimplementedService) Middlewares() []gin.HandlerFunc {
	return nil
}

func (service *UnimplementedService) RegisterHandler(*Engine, *gin.RouterGroup) error {
	return errors.New("UnimplementedServiceRegisterHandler")
}
