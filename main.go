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
	smsClient := nexmo.NewNexmoSMSClient(apiKey, apiSecret)
	smsClient.Send()

	fmt.Println("End")
}
