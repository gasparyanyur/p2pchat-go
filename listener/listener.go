package listener

type OnListener interface {
	// receiving subscribe request if user want to message with you
	OnSubscribe(subscriberIp, subscriberName, subscriberId, subscriberPubKey string)
	// receive applying subscription after sending a subscribe request to user
	//onAllowSubscribe(subscriberName, subscriberId, ownPublicKey string)

	// receiving message from already subscribed users
	OnMessage(username, message, signature string, timStamp uint32)

	// getting error when error caused during process
	OnError(string)
}
