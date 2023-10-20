package alihouse

// AlibabaAlihouseExistinghomeHouseTradeQueryStatusResult 结构体
type AlibabaAlihouseExistinghomeHouseTradeQueryStatusResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败成功信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果对象
	Data *SyncHouseStatusDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
