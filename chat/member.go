package chat

import (
	"crypto/rsa"
)

type Person struct {
	Name      string
	IpAddress string
	PeerId    string
}

type Member struct {
	Person
	PublicKey *rsa.PublicKey
}

type Owner struct {
	Person
	PrivateKey *rsa.PrivateKey
}
