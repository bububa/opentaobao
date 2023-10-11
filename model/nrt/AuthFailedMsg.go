package nrt

// AuthFailedMsg 结构体
type AuthFailedMsg struct {
	// 入参手机号
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 添加失败原因
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
}
