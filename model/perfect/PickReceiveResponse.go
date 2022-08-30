package perfect

// PickReceiveResponse 结构体
type PickReceiveResponse struct {
	// 1
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 1
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 1
	Success string `json:"success,omitempty" xml:"success,omitempty"`
}
