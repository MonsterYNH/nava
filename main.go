package main

import (
	"github.com/MonsterYNH/nava/server"
	"github.com/MonsterYNH/nava/services/helloworld"
)

func main() {
	server.RegisterAPIV1Service(new(helloworld.HelloworldServiceEntry))
	if err := server.Run(); err != nil {
		panic("[ERROR] run app failed, error:" + err.Error())
	}
}
