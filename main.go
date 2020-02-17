package main

import (
	"fmt"

	"github.com/lornajane/goclient-lib/nexmo"
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
	smsClient.Send("LornaTest", "44777000777", "This is a message from golang")

	fmt.Println("End")
}
