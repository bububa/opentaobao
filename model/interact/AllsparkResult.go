package interact

// AllsparkResult 结构体
type AllsparkResult struct {
	// 出错提示
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 出错代码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 服务结果对象
	Data *ActivityWriteResult `json:"data,omitempty" xml:"data,omitempty"`
	// 是否注册成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
