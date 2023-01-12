package entity

import "fmt"

type HelloRequest struct {
	Name string
}

type HelloResponse struct {
	Name    string
	Message string
}

func (h *HelloResponse) ToString() string {
	res := fmt.Sprintf("Hello, %s! This is message for you %s", h.Name, h.Message)
	return res
}
