package alihouse

// AlibabaAlihouseStoreCheckResult 结构体
type AlibabaAlihouseStoreCheckResult struct {
	// 结果列表
	Data []CompanyStoreForCheckDto `json:"data,omitempty" xml:"data>company_store_for_check_dto,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
