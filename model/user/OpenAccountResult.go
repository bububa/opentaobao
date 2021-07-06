package user

// OpenAccountResult 结构体
type OpenAccountResult struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// Open Account信息
	Data *OpenAccount `json:"data,omitempty" xml:"data,omitempty"`
	// 错误码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
}
