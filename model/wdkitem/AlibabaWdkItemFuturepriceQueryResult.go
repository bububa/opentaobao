package wdkitem

// AlibabawdkitemfuturepricequeryResult 结构体
type AlibabawdkitemfuturepricequeryResult struct {
	// 商品未来价
	Models []AlibabawdkitemfuturepricequeryModel `json:"models,omitempty" xml:"models>alibabawdkitemfuturepricequery_model,omitempty"`
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 异常信息
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
