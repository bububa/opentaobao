package user

// OpenaccountLong 结构体
type OpenaccountLong struct {
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回数据
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 返回码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
}
