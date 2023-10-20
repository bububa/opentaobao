package wdk

// AlibabawdkitemcurrentpricequeryResult 结构体
type AlibabawdkitemcurrentpricequeryResult struct {
	// 返回的当前当前商品价格列表
	Models []AlibabawdkitemcurrentpricequeryModel `json:"models,omitempty" xml:"models>alibabawdkitemcurrentpricequery_model,omitempty"`
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 异常信息
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
