package user

// MiniAppResult 结构体
type MiniAppResult struct {
	// 错误提示
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 用户数据
	Data *UserInfoDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}
