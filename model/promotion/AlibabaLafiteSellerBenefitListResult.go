package promotion

// AlibabalafitesellerbenefitlistResult 结构体
type AlibabalafitesellerbenefitlistResult struct {
	// 错误信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回查询结果
	Data *Page `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
