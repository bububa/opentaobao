package logistic

// Addressdtos 结构体
type Addressdtos struct {
	// sender address
	Sender *AeopWlDeclareAddressDto `json:"sender,omitempty" xml:"sender,omitempty"`
	// pickup address
	Pickup *AeopWlDeclareAddressDto `json:"pickup,omitempty" xml:"pickup,omitempty"`
	// receiver address
	Receiver *AeopWlDeclareAddressDto `json:"receiver,omitempty" xml:"receiver,omitempty"`
	// refund address
	Refund *AeopWlDeclareAddressDto `json:"refund,omitempty" xml:"refund,omitempty"`
}
