package main

import (
	"fmt"

	"github.com/antihax/optional"
	"github.com/lornajane/goclient-lib/sms"
)

func main() {
	fmt.Println("Hello")

	smsClient := sms.NewAPIClient(sms.NewConfiguration())
	smsOpts := sms.SendAnSmsOpts{}
	smsOpts.Text = optional.NewString("Hello world")
	smsOpts.ApiSecret = optional.NewString("API_SECRET")

	result, _, err := smsClient.DefaultApi.SendAnSms(nil, "json", "API_KEY", "GoTesting", "NUMBER", &smsOpts)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", result)

	fmt.Println("End")
}
