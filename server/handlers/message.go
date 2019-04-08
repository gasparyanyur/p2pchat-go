package handlers

import (
	"context"
	"github.com/evenfound/even-p2pchat-go/server/api"
)

type MessageHandler struct {
	errorHandler chan<- error
}

func (handler *MessageHandler) ReceiveMessage(ctx context.Context, request *api.MessageRequest) (*api.MessageResponse, error) {

	return &api.MessageResponse{}, nil
}

func NewMessageHandler(errorHandler chan<- error) *MessageHandler {
	return &MessageHandler{
		errorHandler: errorHandler,
	}
}
