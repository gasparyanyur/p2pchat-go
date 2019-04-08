package chat

import (
	"errors"
	"github.com/evenfound/even-crypto"
	"github.com/evenfound/even-p2pchat-go/server/api"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

type Chat struct {
	id string

	member *Member
	owner  *Owner

	client *grpc.ClientConn
}

// Creating a new chat room
func NewChatRoom(member *Member, owner *Owner) (*Chat, error) {
	var conn, err = grpc.Dial(member.IpAddress, grpc.WithInsecure())

	if err != nil {
		return &Chat{}, err
	}

	var chatId = uuid.New()

	var room = &Chat{
		id:     chatId.String(),
		member: member,
		owner:  owner,

		client: conn,
	}

	if room.member.PublicKey == nil {
		var pubKey = room.owner.PrivateKey.PublicKey
		var publicKeyString = crypto.RSAPublicToString(&pubKey)
		err := room.sendPermissionRequest(publicKeyString)

		if err != nil {
			return nil, err
		}
	}

	return room, nil
}

func (chat *Chat) sendPermissionRequest(publicKey string) error {
	client := api.NewSubscribeServiceClient(chat.client)

	var ctx = context.Background()

	clientDeadline := time.Now().Add(time.Duration(5000) * time.Millisecond)

	ctx, _ = context.WithDeadline(ctx, clientDeadline)

	response, err := client.Receive(ctx, &api.SubscribeRequest{
		Username:  chat.owner.Name,
		PublicKey: publicKey,
		Status:    api.SubscribeRequest_Send,
	})

	if err != nil {
		return nil
	}

	if !response.Status {
		return errors.New(response.Message)
	}

	if err != nil {
		return err
	}
	return nil
}
