package main

import (
	"fmt"

	"github.com/iamnbd2022/ianmodules/entry"
	"github.com/iamnbd2022/ianmodules/handler"
)

type Test0Handler struct {
	processor *handler.Processor
}

func (th *Test0Handler) register(processor *handler.Processor) error {
	fmt.Printf("Starting to register handlers \n")
	th.processor = processor
	processor.RegisterRequestHandler(&handler.Matcher{Path: "_health_", Method: "*"}, handler.HealthHanlder)
	processor.RegisterRequestHandler(&handler.Matcher{Path: "hello", Method: "*"}, func(context *handler.IanContext, req *handler.IanRequest) (*handler.IanResponse, error) {
		return &handler.IanResponse{
			Body: "ian"}, nil
	})
	return nil
}

func main() {
	fmt.Printf("Hello ian!")
	handler := Test0Handler{}
	entry.Start("", handler.register, true, nil)
}
