package alilabs

// DeviceStatusDto 结构体
type DeviceStatusDto struct {
	// payload
	Payload *Payload `json:"payload,omitempty" xml:"payload,omitempty"`
	// header
	Header *IotCommonHeader `json:"header,omitempty" xml:"header,omitempty"`
}
