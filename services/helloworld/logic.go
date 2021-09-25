package helloworld

import (
	"fmt"

	"github.com/MonsterYNH/nava/engine"
)

type HelloworldServiceLogic struct{}

func (logic *HelloworldServiceLogic) SayHello(request *HelloworldRequest) *engine.APIResponse {
	return &engine.APIResponse{
		ErrCode: engine.ErrSuccess,
		Data: HelloworldResponse{
			Greating: fmt.Sprintf("Hello %s", request.Name),
		},
	}
}
