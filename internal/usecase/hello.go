package usecase

import (
	"context"

	"github.com/tokopedia/repo-pattern-ex/internal/entity"
	"github.com/tokopedia/repo-pattern-ex/internal/repo"
)

type HelloResource interface {
	Say(ctx context.Context, request entity.HelloRequest) string
}
type HelloUsecase struct {
	Repository repo.HelloRepo
}

func New(helloRepo repo.HelloRepo) HelloResource {
	return &HelloUsecase{
		Repository: helloRepo,
	}
}

func (h *HelloUsecase) Say(ctx context.Context, request entity.HelloRequest) string {
	result := h.Repository.Say(request)
	return result
}
