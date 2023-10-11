package ascp

// DeliverySendAblitiyAsyncResponse 结构体
type DeliverySendAblitiyAsyncResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 处理结果详细信息
	Data *DataDetails `json:"data,omitempty" xml:"data,omitempty"`
	// 系统成功失败   true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
