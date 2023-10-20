package aliqin

// AlibabaaliqinfciotcardofferResult 结构体
type AlibabaaliqinfciotcardofferResult struct {
	// 结果对象
	Models []AlibabaaliqinfciotcardofferModel `json:"models,omitempty" xml:"models>alibabaaliqinfciotcardoffer_model,omitempty"`
	// 1.成功；2.失败
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 状态
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
