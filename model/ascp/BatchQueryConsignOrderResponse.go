package ascp

// BatchQueryConsignOrderResponse 结构体
type BatchQueryConsignOrderResponse struct {
	// 返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 分页数据
	Data *PageData `json:"data,omitempty" xml:"data,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
