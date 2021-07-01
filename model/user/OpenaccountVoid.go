package user

// OpenaccountVoid 结构体
type OpenaccountVoid struct {
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
	// 错误码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}
