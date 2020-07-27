package api

import (
	"context"

	pb_api "github.com/inari111/layered-architecture-example-2020/rpc/api"
)

type taskService struct {
}

func NewTaskService() pb_api.TaskService {
	return &taskService{}
}

func (t *taskService) Create(ctx context.Context, request *pb_api.TaskCreateRequest) (*pb_api.TaskCreateResponse, error) {
	return nil, nil
}
