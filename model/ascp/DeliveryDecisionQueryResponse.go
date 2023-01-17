package ascp

// DeliveryDecisionQueryResponse 结构体
type DeliveryDecisionQueryResponse struct {
	// 详细信息
	Data []DataDetail `json:"data,omitempty" xml:"data>data_detail,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 0=全部失败，1=全部成功，2=部分成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 系统成功失败   true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
