package main

import (
	"bytes"
	"crypto"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
)

func Upload(target string, jwt string) {
	url := getPostUrl(target)
	if url == ""{
		panic("Could't find post url")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jwt)))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func getPostUrl(target string) string {
	records, _ := net.LookupTXT(target)

	fmt.Println("Records for ",target, "are", records)

	var record = getTxtRecord(target, postUrlIdentifier)

	if record == "" {
		panic("TXT Record 'gocom target:' missing")
	}else{
		fmt.Println("TXT Record 'gocom target:' is", record)
	}

	if _, err := url.ParseRequestURI(record); err != nil {
		panic(err)
	}else{
		return record
	}
}

func getTxtRecord(host string, prefix string) string{
	records, _ := net.LookupTXT(host)
	for _, v := range records {
		if strings.HasPrefix(v, prefix) {
			fmt.Println("Found", prefix,"with", v)
			splits :=strings.Split(v, prefix)

			return strings.TrimSpace(splits[len(splits)-1])
		}
	}

	return ""
}

func getPublicKey(target string) crypto.PublicKey {
	base64PublicKey := getTxtRecord(target, publicKeyIdentifier)

	if base64PublicKey == "" {
		panic("No public key")
	}

	data, _ := base64.StdEncoding.DecodeString(base64PublicKey)
	pub, _ := x509.ParsePKIXPublicKey(data)

	return pub
}
