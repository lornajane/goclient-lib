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

import (
	"bytes"
	"encoding/json"
)

// DeliveryReceipt struct for DeliveryReceipt
type DeliveryReceipt struct {
	// The number the message was sent to. Numbers are specified in E.164 format.
	Msisdn *string `json:"msisdn,omitempty"`
	// The SenderID you set in `from` in your request.
	To *string `json:"to,omitempty"`
	// The Mobile Country Code Mobile Network Code (MCCMNC) of the carrier this phone number is registered with.
	NetworkCode *string `json:"network-code,omitempty"`
	// The Nexmo ID for this message.
	MessageId *string `json:"messageId,omitempty"`
	// The cost of the message
	Price *string `json:"price,omitempty"`
	// A code that explains where the message is in the delivery process.
	Status *string `json:"status,omitempty"`
	// When the DLR was recieved from the carrier in the following format `YYMMDDHHMM`. For example, `2001011400` is at `2020-01-01 14:00`
	Scts *string `json:"scts,omitempty"`
	// The status of the request. Will be a non `0` value if there has been an error. See the [Delivery Receipt documentation](https://developer.nexmo.com/messaging/sms/guides/delivery-receipts#dlr-error-codes) for more details
	ErrCode *string `json:"err-code,omitempty"`
	// The time when Nexmo started to push this Delivery Receipt to your webhook endpoint.
	MessageTimestamp *string `json:"message-timestamp,omitempty"`
}

// GetMsisdn returns the Msisdn field value if set, zero value otherwise.
func (o *DeliveryReceipt) GetMsisdn() string {
	if o == nil || o.Msisdn == nil {
		var ret string
		return ret
	}
	return *o.Msisdn
}

// GetMsisdnOk returns a tuple with the Msisdn field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReceipt) GetMsisdnOk() (string, bool) {
	if o == nil || o.Msisdn == nil {
		var ret string
		return ret, false
	}
	return *o.Msisdn, true
}

// HasMsisdn returns a boolean if a field has been set.
func (o *DeliveryReceipt) HasMsisdn() bool {
	if o != nil && o.Msisdn != nil {
		return true
	}

	return false
}

// SetMsisdn gets a reference to the given string and assigns it to the Msisdn field.
func (o *DeliveryReceipt) SetMsisdn(v string) {
	o.Msisdn = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *DeliveryReceipt) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReceipt) GetToOk() (string, bool) {
	if o == nil || o.To == nil {
		var ret string
		return ret, false
	}
	return *o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *DeliveryReceipt) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *DeliveryReceipt) SetTo(v string) {
	o.To = &v
}

// GetNetworkCode returns the NetworkCode field value if set, zero value otherwise.
func (o *DeliveryReceipt) GetNetworkCode() string {
	if o == nil || o.NetworkCode == nil {
		var ret string
		return ret
	}
	return *o.NetworkCode
}

// GetNetworkCodeOk returns a tuple with the NetworkCode field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReceipt) GetNetworkCodeOk() (string, bool) {
	if o == nil || o.NetworkCode == nil {
		var ret string
		return ret, false
	}
	return *o.NetworkCode, true
}

// HasNetworkCode returns a boolean if a field has been set.
func (o *DeliveryReceipt) HasNetworkCode() bool {
	if o != nil && o.NetworkCode != nil {
		return true
	}

	return false
}

// SetNetworkCode gets a reference to the given string and assigns it to the NetworkCode field.
func (o *DeliveryReceipt) SetNetworkCode(v string) {
	o.NetworkCode = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *DeliveryReceipt) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReceipt) GetMessageIdOk() (string, bool) {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret, false
	}
	return *o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *DeliveryReceipt) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *DeliveryReceipt) SetMessageId(v string) {
	o.MessageId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *DeliveryReceipt) GetPrice() string {
	if o == nil || o.Price == nil {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReceipt) GetPriceOk() (string, bool) {
	if o == nil || o.Price == nil {
		var ret string
		return ret, false
	}
	return *o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *DeliveryReceipt) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *DeliveryReceipt) SetPrice(v string) {
	o.Price = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeliveryReceipt) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReceipt) GetStatusOk() (string, bool) {
	if o == nil || o.Status == nil {
		var ret string
		return ret, false
	}
	return *o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeliveryReceipt) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeliveryReceipt) SetStatus(v string) {
	o.Status = &v
}

// GetScts returns the Scts field value if set, zero value otherwise.
func (o *DeliveryReceipt) GetScts() string {
	if o == nil || o.Scts == nil {
		var ret string
		return ret
	}
	return *o.Scts
}

// GetSctsOk returns a tuple with the Scts field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReceipt) GetSctsOk() (string, bool) {
	if o == nil || o.Scts == nil {
		var ret string
		return ret, false
	}
	return *o.Scts, true
}

// HasScts returns a boolean if a field has been set.
func (o *DeliveryReceipt) HasScts() bool {
	if o != nil && o.Scts != nil {
		return true
	}

	return false
}

// SetScts gets a reference to the given string and assigns it to the Scts field.
func (o *DeliveryReceipt) SetScts(v string) {
	o.Scts = &v
}

// GetErrCode returns the ErrCode field value if set, zero value otherwise.
func (o *DeliveryReceipt) GetErrCode() string {
	if o == nil || o.ErrCode == nil {
		var ret string
		return ret
	}
	return *o.ErrCode
}

// GetErrCodeOk returns a tuple with the ErrCode field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReceipt) GetErrCodeOk() (string, bool) {
	if o == nil || o.ErrCode == nil {
		var ret string
		return ret, false
	}
	return *o.ErrCode, true
}

// HasErrCode returns a boolean if a field has been set.
func (o *DeliveryReceipt) HasErrCode() bool {
	if o != nil && o.ErrCode != nil {
		return true
	}

	return false
}

// SetErrCode gets a reference to the given string and assigns it to the ErrCode field.
func (o *DeliveryReceipt) SetErrCode(v string) {
	o.ErrCode = &v
}

// GetMessageTimestamp returns the MessageTimestamp field value if set, zero value otherwise.
func (o *DeliveryReceipt) GetMessageTimestamp() string {
	if o == nil || o.MessageTimestamp == nil {
		var ret string
		return ret
	}
	return *o.MessageTimestamp
}

// GetMessageTimestampOk returns a tuple with the MessageTimestamp field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryReceipt) GetMessageTimestampOk() (string, bool) {
	if o == nil || o.MessageTimestamp == nil {
		var ret string
		return ret, false
	}
	return *o.MessageTimestamp, true
}

// HasMessageTimestamp returns a boolean if a field has been set.
func (o *DeliveryReceipt) HasMessageTimestamp() bool {
	if o != nil && o.MessageTimestamp != nil {
		return true
	}

	return false
}

// SetMessageTimestamp gets a reference to the given string and assigns it to the MessageTimestamp field.
func (o *DeliveryReceipt) SetMessageTimestamp(v string) {
	o.MessageTimestamp = &v
}

type NullableDeliveryReceipt struct {
	Value DeliveryReceipt
	ExplicitNull bool
}

func (v NullableDeliveryReceipt) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}	
}

func (v *NullableDeliveryReceipt) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
