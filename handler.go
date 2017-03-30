package main

import (
	"log"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

type CustomError struct {
	s string
}

func (e *CustomError) Error() string {
	return e.s
}

func Handle(evt interface{}, ctx *runtime.Context) (interface{}, error) {
	log.Println("Hello, World Logging!")
	return "Hello, World!", nil
	// return nil, &CustomError{"somthing bad happened"}
}
