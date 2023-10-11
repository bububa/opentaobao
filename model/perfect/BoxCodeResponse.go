package perfect

// BoxCodeResponse 结构体
type BoxCodeResponse struct {
	// 1
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 1
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 箱号
	BoxCode string `json:"box_code,omitempty" xml:"box_code,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
