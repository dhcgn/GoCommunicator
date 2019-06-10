package main

import (
	"encoding/json"
	"io/ioutil"
)

type Message struct {
	Host string
	Message string
	Uid string
}

func uploadMessage(msg string, uid string, host string) {
	m := Message{host, msg, uid}
	b, _ := json.Marshal(m)
	jwt := encrypt(b)
	Upload(*target,jwt)
}

func uploadFile() {
	dat, _ := ioutil.ReadFile(*uploadFilePath)
	jwt := encrypt(dat)
	Upload(*target, jwt)
}
