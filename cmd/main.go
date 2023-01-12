package main

import (
	"net/http"

	"github.com/tokopedia/repo-pattern-ex/internal/handler"
	"github.com/tokopedia/repo-pattern-ex/internal/repo"
	"github.com/tokopedia/repo-pattern-ex/internal/usecase"
)

func main() {
	repo := repo.New()
	usecaseHello := usecase.New(*repo)
	handler := handler.New(usecaseHello)

	// generic var-example
	var w http.ResponseWriter
	var r http.Request
	// can invoke to this
	handler.Say(w, &r)
}
