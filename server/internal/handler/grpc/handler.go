package grpc

import (
	"context"

	"github.com/dnguyen316/GoGetIt/server/internal/generated/grpc/go_get_it"
	"github.com/dnguyen316/GoGetIt/server/internal/logic"
)

type Handler struct {
	go_get_it.GoGetItServiceServer
	accountLogic logic.Account
}

func NewHandler(
	accountLogic logic.Account,
) go_get_it.GoGetItServiceServer {
	return &Handler{
		accountLogic: accountLogic,
	}
}

func (a *Handler) CreateAccount(ctx context.Context, request *go_get_it.CreateAccountRequest) (*go_get_it.CreateAccountResponse, error) {
	output, err := a.accountLogic.CreateAccount(ctx, logic.CreateAccountParams{
		AccountName: request.GetAccountName(),
		Password:    request.GetPassword(),
	})

	if err != nil {
		return nil, err
	}

	return &go_get_it.CreateAccountResponse{
		AccountId: output.ID,
	}, nil
}

// CreateSession handles creation of a new session.
func (h *Handler) CreateSession(ctx context.Context, req *go_get_it.CreateSessionRequest) (*go_get_it.CreateSessionResponse, error) {
	// TODO: Implement business logic
	return &go_get_it.CreateSessionResponse{}, nil
}

// CreateDownloadTask handles the creation of a new download task.
func (h *Handler) CreateDownloadTask(ctx context.Context, req *go_get_it.CreateDownloadTaskRequest) (*go_get_it.CreateDownloadTaskResponse, error) {
	// TODO: Implement business logic
	return &go_get_it.CreateDownloadTaskResponse{}, nil
}

// GetDownloadTaskList retrieves the list of download tasks for the accountunt.
func (h *Handler) GetDownloadTaskList(ctx context.Context, req *go_get_it.GetDownloadTaskListRequest) (*go_get_it.GetDownloadTaskListResponse, error) {
	// TODO: Implement business logic
	return &go_get_it.GetDownloadTaskListResponse{}, nil
}

// UpdateDownloadTask handles updates to an existing download task.
func (h *Handler) UpdateDownloadTask(ctx context.Context, req *go_get_it.UpdateDownloadTaskRequest) (*go_get_it.UpdateDownloadTaskResponse, error) {
	// TODO: Implement business logic
	return &go_get_it.UpdateDownloadTaskResponse{}, nil
}

// DeleteDownloadTask handles deletion of a specific download task.
func (h *Handler) DeleteDownloadTask(ctx context.Context, req *go_get_it.DeleteDownloadTaskRequest) (*go_get_it.DeleteDownloadTaskResponse, error) {
	// TODO: Implement business logic
	return &go_get_it.DeleteDownloadTaskResponse{}, nil
}

// GetDownloadTaskFile streams back the file associated with a download task.
// This example simply returns nil without streaming any data.
func (h *Handler) GetDownloadTaskFile(req *go_get_it.GetDownloadTaskFileRequest, stream go_get_it.GoGetItService_GetDownloadTaskFileServer) error {
	// TODO: Implement streaming logic (e.g., sending chunks of file data)
	return nil
}
