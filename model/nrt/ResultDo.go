package nrt

// ResultDo 结构体
type ResultDo struct {
	// 系统自动生成
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 系统自动生成
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误码
	Errcode string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 异常信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 参数对象
	Data *NrtStoreDto `json:"data,omitempty" xml:"data,omitempty"`
	// 调用是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}
