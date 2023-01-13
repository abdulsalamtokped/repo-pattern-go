package repo

import (
	"github.com/tokopedia/repo-pattern-ex/internal/entity"
)

type HelloRepo struct{}

func New() *HelloRepo {
	return &HelloRepo{}
}

func (h *HelloRepo) Say(request entity.HelloRequest) string {
	var helloEntity entity.HelloResponse

	helloEntity.Name = request.Name
	helloEntity.Message = "it's works!"

	return helloEntity.ToString()
}
