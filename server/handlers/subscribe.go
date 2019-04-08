package handlers

import (
	"context"
	"errors"
	"github.com/evenfound/even-p2pchat-go/listener"
	"github.com/evenfound/even-p2pchat-go/server/api"
	"google.golang.org/grpc/peer"
)

const (
	SubscribeResponseWaiting = "please wait for member verifying"

	SubscribeIncorrectData = "can not get necessary information about subscriber"
)

type SubscribeHandler struct {
	errorHandler chan<- error
	listener     listener.OnListener
}

func NewSubscribeHandler(errorHandler chan<- error, listener listener.OnListener) *SubscribeHandler {
	return &SubscribeHandler{
		errorHandler: errorHandler,
		listener:     listener,
	}
}

func (handler *SubscribeHandler) Receive(ctx context.Context, request *api.SubscribeRequest) (*api.SubscribeResponse, error) {
	var client, ok = peer.FromContext(ctx)

	if !ok {
		return &api.SubscribeResponse{
			Status:  false,
			Message: SubscribeIncorrectData,
		}, errors.New(SubscribeIncorrectData)
	}

	var subscriberIpAddress = client.Addr.String()

	handler.listener.OnSubscribe(subscriberIpAddress, request.Username, "", request.PublicKey)

	return &api.SubscribeResponse{
		Status: true,
	}, nil
}

func (handler *SubscribeHandler) Apply(ctx context.Context, request *api.AllowSubscribeMessage) (*api.SubscribeResponse, error) {

	return nil, nil
}
