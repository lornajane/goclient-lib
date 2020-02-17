package nexmo

import (
	"fmt"

	"github.com/antihax/optional"
	"github.com/lornajane/goclient-lib/sms"
)

type NexmoSMSClient struct {
	apiKey    string
	apiSecret string
}

func NewNexmoSMSClient(Auth Auth) *NexmoSMSClient {
	client := new(NexmoSMSClient)
	creds := Auth.getCreds()
	client.apiKey = creds[0]
	client.apiSecret = creds[1]
	return client
}

func (client *NexmoSMSClient) Send() {

	config := sms.NewConfiguration()

	// for local Prism testing
	config.BasePath = "http://localhost:4010"

	smsClient := sms.NewAPIClient(config)

	smsOpts := sms.SendAnSmsOpts{}
	smsOpts.Text = optional.NewString("Hello world")
	smsOpts.ApiSecret = optional.NewString(client.apiSecret)

	result, _, err := smsClient.DefaultApi.SendAnSms(nil, "json", client.apiKey, "GoTesting", "447846810475", &smsOpts)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", result)

	msgs := result.GetMessages()
	fmt.Println(msgs[0].GetRemainingBalance())

}
