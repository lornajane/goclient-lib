package main

import (
	"fmt"

	"github.com/lornajane/goclient-lib/nexmo"
	"github.com/lornajane/goclient-lib/sms"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Hello")

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	apiKey := viper.GetString("api_key")
	apiSecret := viper.GetString("api_secret")

	auth := nexmo.CreateAuthFromKeySecret(apiKey, apiSecret)
	smsClient := nexmo.NewNexmoSMSClient(auth)

	smsConfig := sms.NewConfiguration()
	// for local Prism testing
	smsConfig.BasePath = "http://localhost:4010?__example=rhubarb"

	response, err := smsClient.Send("LornaTest", "44777000777", "This is a message from golang", nexmo.SMSClientOpts{Config: smsConfig})
	// response, err := smsClient.Send("LornaTest", "44777000777", "This is a message from golang", nexmo.SMSClientOpts{})
	fmt.Printf("%#v\n", response)

	if err != nil {
		panic(err)
	}

	if response.Messages[0].Status == "0" {
		fmt.Println("Account Balance: " + response.Messages[0].RemainingBalance)
	}

	fmt.Println("End")
}
