package promotion

// AlibabalafiteselleractivitylistResult 结构体
type AlibabalafiteselleractivitylistResult struct {
	// 错误描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 接口返回结果
	PageData *Page `json:"page_data,omitempty" xml:"page_data,omitempty"`
}
