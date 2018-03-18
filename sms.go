package main

import (
	"net/url"
	"net/http"
	"strings"
	"fmt"
	"encoding/json"
	"os"
)

var (
	BaseUrl string = "https://api.twilio.com/2010-04-01"
	AccountSid = os.Getenv("ACCOUNT_SID")
	AuthToken = os.Getenv("AUTH_TOKEN")
)

func genMessageReader() strings.Reader {
	message := url.Values{}
	message.Set("TO", "+")
	message.Set("FROM", "+")
	message.Set("Body", "Initial message!")
	reader :=  *strings.NewReader(message.Encode())
	return reader
}

func SendSms() {
	client := &http.Client{}
	reader := genMessageReader()
	smsEndpoint := fmt.Sprintf("%s/Accounts/%s/Messages", BaseUrl, AccountSid)
	fmt.Printf("smsEndpoint: %s\n", smsEndpoint)
	//res, _ := client.PostForm(smsEndpoint, reader)

	req, _ := http.NewRequest("POST", smsEndpoint, &reader)
	req.SetBasicAuth(AccountSid, AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := client.Do(req)
	fmt.Println()
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		var data map[string]interface{}
		err := json.NewDecoder(res.Body).Decode(&data)
		if err == nil {
			fmt.Println("No error found")
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Printf("Something went wrong! Code: %d\n", res.StatusCode)
		fmt.Println(res.Status)
	}
}

func main() {
	fmt.Println("SMS!")
	SendSms()
}