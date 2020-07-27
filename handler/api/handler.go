package api

import (
	"context"
	"log"
	"net/http"

	"github.com/twitchtv/twirp"

	pb_api "github.com/inari111/layered-architecture-example-2020/rpc/api"
)

type server struct {
	pathPrefix string
	twServer   pb_api.TwirpServer
}

func NewHandler(
	taskService pb_api.TaskService,
) http.Handler {
	mux := http.NewServeMux()

	hook := &twirp.ServerHooks{
		RequestReceived:  nil,
		RequestRouted:    nil,
		ResponsePrepared: nil,
		ResponseSent:     nil,
		Error:            errorHook,
	}

	hooks := twirp.ChainHooks(hook)

	var servers = []server{
		{
			pathPrefix: pb_api.TaskServicePathPrefix,
			twServer: pb_api.NewTaskServiceServer(
				taskService,
				hooks,
			),
		},
	}

	for _, s := range servers {
		mux.Handle(s.pathPrefix, s.twServer)
	}

	return mux
}

func errorHook(ctx context.Context, err twirp.Error) context.Context {
	if err != nil {
		log.Printf("%+v", err)
	}
	return ctx
}
