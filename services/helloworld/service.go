package helloworld

import (
	"github.com/MonsterYNH/nava/engine"

	"github.com/gin-gonic/gin"
)

type HelloworldService interface {
	SayHello(*HelloworldRequest) *engine.APIResponse
}

type HelloworldServiceEntry struct {
	engine.UnimplementedService
	HelloworldServiceLogic
}

func (service *HelloworldServiceEntry) Name() string {
	return "helloworld"
}

func (service *HelloworldServiceEntry) RegisterHandler(app *engine.Engine, group *gin.RouterGroup) error {
	group.GET("/sayhello", service.SayHelloHandler)
	return nil
}

func (service *HelloworldServiceEntry) SayHelloHandler(c *gin.Context) {
	var (
		request  = HelloworldRequest{}
		response *engine.APIResponse
	)
	defer func() {
		c.Set(engine.ResponseData, response)
	}()

	if err := c.ShouldBindQuery(&request); err != nil {
		response.ErrCode = engine.ErrParam.SetErrors(err.Error())
		return
	}

	response = service.SayHello(&request)
}

type HelloworldRequest struct {
	Name string `form:"name"`
}

type HelloworldResponse struct {
	Greating string `json:"greating"`
}
