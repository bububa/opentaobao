package nrt

// ResultDo 结构体
type ResultDo struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误码
	Errcode string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 异常信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 数据对象
	Data *NrtEaLoginDto `json:"data,omitempty" xml:"data,omitempty"`
	// 请求结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 调用是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}
