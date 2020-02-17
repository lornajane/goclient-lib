/*
 * SMS API
 *
 * With the Nexmo SMS API you can send SMS from your account and lookup messages both messages that you've sent as well as messages sent to your virtual numbers. Numbers are specified in E.164 format. More SMS API documentation is at <https://developer.nexmo.com/messaging/sms/overview>
 *
 * API version: 1.0.5
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sms
// Sms struct for Sms
type Sms struct {
	// The amount of messages in the request
	MessageCount string `json:"message-count,omitempty"`
	Messages []Message `json:"messages,omitempty"`
}
