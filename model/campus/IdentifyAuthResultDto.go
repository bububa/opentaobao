package campus

// IdentifyAuthResultDto 结构体
type IdentifyAuthResultDto struct {
	// 业务id
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 鉴权结果码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 鉴权结果消息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 用户信息
	User *UserDto `json:"user,omitempty" xml:"user,omitempty"`
}
