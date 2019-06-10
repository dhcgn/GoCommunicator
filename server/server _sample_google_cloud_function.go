// Package p contains an HTTP Cloud Function.
package p

import (
	"bytes"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	. "gopkg.in/square/go-jose.v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Starting")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, "Read Body - NOK")
	}else{
		fmt.Fprintln(w, "Read Body - OK")
		fmt.Fprintln(w, "Body as string: ", "'"+string(bodyBytes)+"'")

	}

	privateKeyBase64 := "[My Private Key]"
	data, _ := base64.StdEncoding.DecodeString(privateKeyBase64)
	privateKey, err := x509.ParseECPrivateKey(data)
	if err != nil {
		fmt.Fprintln(w, "ParseECPrivateKey - NOK")
		log.Println(err)
	}else{
		fmt.Fprintln(w, "ParseECPrivateKey - OK")
	}

	object, err := ParseEncrypted(string(bodyBytes))
	if err != nil {
		fmt.Fprintln(w, "ParseEncrypted - NOK")
		log.Println(err)
	}else{
		fmt.Fprintln(w, "ParseEncrypted - OK")
		fmt.Fprintln(w, "Algorithm:", string(object.Header.Algorithm))
		fmt.Fprintln(w, "AuthData:", string(object.GetAuthData()))
	}

	decrypted, err := object.Decrypt(privateKey)
	if err != nil {
		fmt.Fprintln(w, "Decrypt - NOK")
		fmt.Fprintln(w,err.Error())
		log.Println(err)
	}else{
		fmt.Fprintln(w, "Decrypt - OK")

	}

	var m Message
	json.Unmarshal(decrypted, &m)


	// fmt.Fprint(w, "Your Message:", m)
	sendPush(m)
}

type Message struct {
	Host string
	Message string
	Uid string
}

func sendPush(msg Message){
	v := url.Values{}
	v.Set("token", "[API Key]")
	v.Add("user", "[User Key]")


	v.Add("title", "Message from: "+msg.Host)
	displayMsg  :=   msg.Message + "\n\n" +
					"UID: "+msg.Uid

	v.Add("message", displayMsg)

	postString := v.Encode()

	url := "http://api.pushover.net/1/messages.json"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(postString)))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}