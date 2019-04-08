package messenger

import (
	"errors"
	"fmt"
	"github.com/evenfound/even-crypto"
	"github.com/evenfound/even-p2pchat-go/chat"
	"github.com/evenfound/even-p2pchat-go/listener"
	"github.com/evenfound/even-p2pchat-go/server"
)

const (
	ProcessNotStarted = "chat service does not started"
)

type P2PMessenger struct {
	errorHandler chan error

	chats map[string]*chat.Chat

	listener *listener.OnListener
	server   *server.Server

	owner *chat.Owner

	started bool
}

// Starting and initializing chat with by passing listener , own privateKey , peer id and username
// Start function starts a new grpc server which will receive all requests like subscription and messages
func (messenger *P2PMessenger) Start(onListener listener.OnListener) {
	go func(messenger *P2PMessenger) {
		var (
			// error handler which will track all process to
			// receive and pass errors during the process
			errorHandler = make(chan error)

			grpcServer = &server.Server{}
		)

		messenger.errorHandler = errorHandler

		// starting new grpc server
		go grpcServer.Start(messenger.errorHandler, onListener)

		go func() {
			for {
				select {
				// if error occurred the process will pass that error to listener
				case err := <-errorHandler:
					onListener.OnError(err.Error())
				}
			}
		}()

		messenger.started = true

		messenger.chats = map[string]*chat.Chat{}

	}(messenger)

}

// checking if the main process is started
// most of the functionality works only started mod
func (messenger *P2PMessenger) checkProcess() bool {
	return messenger.started
}

// @TODO make an authorization by passing owner data (if flatter has permission to sub-package files than pass Owner structure right away)
func (messenger *P2PMessenger) Authorize(ownerId, ownerIpAddress, ownerName, privateKey string) bool {

	// owner is the owner app. every time before start
	// chat you should give information data about owner
	var owner = &chat.Owner{
		Person: chat.Person{
			IpAddress: ownerIpAddress,
			Name:      ownerName,
			PeerId:    ownerId,
		},
	}

	// validation of user private key
	var pkHash, err = crypto.ToRSAPrivateKey(privateKey)

	if err != nil {
		return false
	}
	owner.PrivateKey = pkHash

	messenger.owner = owner

	return true
}

// StartChat starts a new or imports already existing chat room
// If the chat room is a new (when missing alien public key)
// automatically will sent a subscribe message
// @TODO if flatter has permission to sub-package files than you can refactor this code to make function that receive object of Member
func (messenger *P2PMessenger) StartChat(memberIp, memberId, memberName, publicKey string) bool {
	var member = &chat.Member{
		Person: chat.Person{
			IpAddress: memberIp,
			Name:      memberName,
			PeerId:    memberId,
		},
	}

	if publicKey == "" {
		var pkHash, err = crypto.ToRSAPublicKey(publicKey)

		if err != nil {
			messenger.errorHandler <- err
			return false
		}

		member.PublicKey = pkHash
	}

	room, err := chat.NewChatRoom(member, messenger.owner)

	if err != nil {
		messenger.errorHandler <- err
		return false
	}

	fmt.Println(memberName, room)

	messenger.chats[memberName] = room

	return true
}

func (messenger *P2PMessenger) ApplySubscription() {
	if !messenger.checkProcess() {
		messenger.errorHandler <- errors.New(ProcessNotStarted)
		return
	}
}

//func main() {
//	var wg sync.WaitGroup
//
//	wg.Add(1)
//
//	var privateKey = "308204a40201000282010100c4c5f626b4ae9c2046689ae3a8bb86b2b154357e62427b21334a3cc5cb60834f068e9f32f7370ba4a05f00afdae77b580621e282147e06794826e6b0d2b4c41f1ec17fe260ff3a7921792568c4205f1c029632638bf343bc05e903467539f9c9b7791b4e55ee11fa816eaa23c80379206f0e965bd69410c6c1e375df14be386675be76994766104013ba18649ebff76f727142006a08c6aea934df9438b451049fb32f77861d66d46a753d68e3b40152d017639241aa0ee6be446c403063d27af64ef1ffb6e3d36ae7cb4a12c8fbb428974968ea218de01a98fd7ab7589fe3296936c9817bcf362b768e79fa770b8737ae221072e5718f6160bcfa64685410750203010001028201004879492214664366b7c80b54526f4f3b3d88f072ee29e243e62a2f9c023e37dec161824d393472fb5d7de038e4fd6136987b9e7b9ea49429d36350904beafc5921f0cef3481a7924e829409807b48c933a78ab7272b754794a0bfa82d6e65593f01990ef17432a2803c808ae491ace601dc757a5d7b08d45476049a7b272253f2215c7ab06b1845544be3764c9c4bcd26b31cee26c66c0c52c133099845bb275006f2691e106be723213d112ad7e1a2362cd2da762ba965db8a8fd3eaefb4045dc61d8e5efa0a057a1f8de68ddc51fc36e4b170a658cf77b86fa5640d4fba9d856188b9ddb8f6ef160b1fb864c666359ac6881529cf85e76d6cefa0e30a8844102818100d5f8891fba58c6c7d3d141f7c601f72b0434cc93490277d9c4232b386e7d1ccb3fb4c1689e69b2e176c9563cee7a5b7510986362e7aa4751034135062e05028a141b7aefe146a3c3914ac739018b38ad0ae9377c3ef53307cd20f6a6fe0ce098ade45d273ea1af51458edd47e74fc08c0f29a30c9296254fc1d66d3b732df53d02818100eb6ca6cfaa515d56215ee09e0095091f13c46c37d518bfd8b99fdb201cb70948ef46692052eb37fad6404e878ae7d180dfec1b6f830b1db891e5bba1a5463c77bd102273b6936874783e747374199b049d980749a65ce437bddd030b3ffdbca1a8b5d2a122d7413a5e19c36d2e0c412fc5366e57a46cf3972cd959492b8c6b9902818100befac1718742fbd199f6d9998cab6e707fa59cbf0d585d9f0a1895d81d8e20282d696983b6790347885db652564cf1dab000441d204774b613609d3231db18e423844f9b520859dd118e9bf8feaec3dc77b3a769874c15cb2dc4f14225d95a920c0b9b358f0a834fd7f4e6fee9afde5d649f363fdf8a922cf90729a1f97118c502818100b2783194e816ba5d1afb23e863a497bf996aa4674702761d43caed80d083e96403102b8db78e5d67a8982370195b57b50d0b9e58d0ccd2812309374e794e5f749e3ce701357d56084547835c2abd6688ff374aff08410f393a939b452203a6c61e7187f563c62ecaa29f8148f9498cf93bd5e19c4b27d519db84016db8b7c0b1028180161121413f69c5733afba93cb7dca4e372ca63a8b178011eb1fcdabeddf2835c7fa9379b21c074e217c65ecaf238d8b268fcd71e6e5bc944e6e83d971ffd44dae3c062aee9e38da4f82ebf586b679e9191348ce57956d9da6d55d870095bb2f2202a540a02e9347d68ab4fce33d7df58f8f7293ddb421c1332363c8bc7e9ed56"
//
//	//var publicKey = "3082010a0282010100c4c5f626b4ae9c2046689ae3a8bb86b2b154357e62427b21334a3cc5cb60834f068e9f32f7370ba4a05f00afdae77b580621e282147e06794826e6b0d2b4c41f1ec17fe260ff3a7921792568c4205f1c029632638bf343bc05e903467539f9c9b7791b4e55ee11fa816eaa23c80379206f0e965bd69410c6c1e375df14be386675be76994766104013ba18649ebff76f727142006a08c6aea934df9438b451049fb32f77861d66d46a753d68e3b40152d017639241aa0ee6be446c403063d27af64ef1ffb6e3d36ae7cb4a12c8fbb428974968ea218de01a98fd7ab7589fe3296936c9817bcf362b768e79fa770b8737ae221072e5718f6160bcfa64685410750203010001"
//
//	var testListener = OutListener{}
//
//	var testMessenger = P2PMessenger{}
//
//	testMessenger.Authorize("GUID-123", "127.0.0.1:8094", "yuri2", privateKey)
//
//	testMessenger.Start(&testListener)
//
//	//testMessenger.StartChat(":8093", "GUID-123", "123")
//
//	wg.Wait()
//
//}
