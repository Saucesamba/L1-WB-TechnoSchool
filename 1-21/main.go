package main

import (
	"fmt"
)

type ClientInterface interface {
	ProcessRequest()
}

type Adapter struct {
	service Service
}

func (a *Adapter) ProcessRequest() {
	fmt.Println("Adapting")
	a.service.ProcessData()
}

type Service struct {
}

func (s *Service) ProcessData() {
	fmt.Println("Request is processing")
}
func SendRequest(client ClientInterface) {
	client.ProcessRequest()
}

func NewAdapter(s Service) *Adapter {
	return &Adapter{service: s}
}

func main() {
	service := new(Service)

	adapter := NewAdapter(*service)

	SendRequest(adapter)
}
