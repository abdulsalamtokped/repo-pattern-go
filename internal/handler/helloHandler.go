package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/tokopedia/repo-pattern-ex/internal/entity"
	"github.com/tokopedia/repo-pattern-ex/internal/usecase"
)

type HelloHandler interface {
	Say(w http.ResponseWriter, r *http.Request) error
}

// Handler defines hello handler
type handler struct {
	Usecase usecase.HelloResource // usecase field, make it interface so it can be injected during test
}

func New(u usecase.HelloResource) HelloHandler {
	return &handler{
		Usecase: u,
	}
}
func (h *handler) Say(w http.ResponseWriter, r *http.Request) error {
	request := entity.HelloRequest{}
	json.NewDecoder(r.Body).Decode(&request)
	result := h.Usecase.Say(context.Background(), request)

	w.Write([]byte(result))
	return nil
}
