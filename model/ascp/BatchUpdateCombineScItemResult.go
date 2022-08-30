package ascp

// BatchUpdateCombineScItemResult 结构体
type BatchUpdateCombineScItemResult struct {
	// 结果明细
	Detail []DetailItem `json:"detail,omitempty" xml:"detail>detail_item,omitempty"`
	// 0=全部失败，1=全部成功，2=部分成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
