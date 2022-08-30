package alihouse

// AlibabaAlihouseExistinghomeEntrustsellingQueryResult 结构体
type AlibabaAlihouseExistinghomeEntrustsellingQueryResult struct {
	// 系统自动生成
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data *CustomerEntrustSellingDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
