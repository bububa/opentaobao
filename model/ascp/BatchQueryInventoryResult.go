package ascp

// BatchQueryInventoryResult 结构体
type BatchQueryInventoryResult struct {
	// 结果明细
	Detail []ScItemInfo `json:"detail,omitempty" xml:"detail>sc_item_info,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
