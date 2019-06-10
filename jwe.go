package main

import (
	. "gopkg.in/square/go-jose.v2"
)

func encrypt(data []byte) string {
	publicKey := getPublicKey(*target)

	encrypter, _ := NewEncrypter(A256GCM, Recipient{Algorithm: ECDH_ES_A256KW, Key: publicKey}, nil)

	// id, _ := machineid.ProtectedID(AppName)
	object, _ := encrypter.Encrypt(data)

	compactSerialized, _ := object.CompactSerialize()

	return compactSerialized
}
