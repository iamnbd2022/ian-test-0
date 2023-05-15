package main

import (
	"fmt"

	"github.com/iamnbd2022/ianmodules/entry"
	"github.com/iamnbd2022/ianmodules/handler"
)

var processor = handler.Processor{}

func register() {
	fmt.Printf("Starting to register handlers \n")
	processor.RegisterRequestHandler(&handler.Matcher{Path: "_health_", Method: "*"}, handler.HealthHanlder)
	processor.RegisterRequestHandler(&handler.Matcher{Path: "hello", Method: "*"}, func(context *handler.IanContext, req *handler.IanRequest) (*handler.IanResponse, error) {
		return &handler.IanResponse{
			Body: "ian"}, nil
	})
}

func main() {
	fmt.Printf("Hello ian!")
	register()
	entry.Start(&processor)
}
