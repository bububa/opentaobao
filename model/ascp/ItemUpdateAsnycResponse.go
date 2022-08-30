package ascp

// ItemUpdateAsnycResponse 结构体
type ItemUpdateAsnycResponse struct {
	// 业务处理结果
	Data []DataItem `json:"data,omitempty" xml:"data>data_item,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 0=全部接收失败，1=全部接收成功，2=部分接受成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
