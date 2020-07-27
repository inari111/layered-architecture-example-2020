// +build wireinject

package di

import (
	"net/http"

	"github.com/google/wire"

	"github.com/inari111/layered-architecture-example-2020/handler/api"
)

func InitializeAPIHandler() http.Handler {
	wire.Build(
		api.NewTaskService,

		api.NewHandler,
	)
	return nil
}
