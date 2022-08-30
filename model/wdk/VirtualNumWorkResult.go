package wdk

// VirtualNumWorkResult 结构体
type VirtualNumWorkResult struct {
	// 淘鲜达交易单号
	SourceOrderId string `json:"source_order_id,omitempty" xml:"source_order_id,omitempty"`
	// 查询类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 返回虚拟号
	VirtualNumber string `json:"virtual_number,omitempty" xml:"virtual_number,omitempty"`
	// 订阅id
	SubId string `json:"sub_id,omitempty" xml:"sub_id,omitempty"`
}
