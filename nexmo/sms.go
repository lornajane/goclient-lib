package nexmo

import (
	"fmt"

	"github.com/antihax/optional"
	"github.com/lornajane/goclient-lib/sms"
)

// Client for working with the SMS API
type NexmoSMSClient struct {
	apiKey    string
	apiSecret string
}

// Create a new SMS Client, supplying an Auth to work with
func NewNexmoSMSClient(Auth Auth) *NexmoSMSClient {
	client := new(NexmoSMSClient)
	creds := Auth.getCreds()
	client.apiKey = creds[0]
	client.apiSecret = creds[1]
	return client
}

// Send an SMS
func (client *NexmoSMSClient) Send(from string, to string, text string) {

	config := sms.NewConfiguration()

	// for local Prism testing
	config.BasePath = "http://localhost:4010"

	smsClient := sms.NewAPIClient(config)

	smsOpts := sms.SendAnSmsOpts{}
	smsOpts.Text = optional.NewString(text)
	smsOpts.ApiSecret = optional.NewString(client.apiSecret)

	result, _, err := smsClient.DefaultApi.SendAnSms(nil, "json", client.apiKey, from, to, &smsOpts)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", result)

	msgs := result.GetMessages()
	fmt.Println(msgs[0].GetRemainingBalance())

}
