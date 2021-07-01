package examination

// Abnormalitemlist 结构体
type Abnormalitemlist struct {
	// 异常项指标名称
	ItemDetail string `json:"item_detail,omitempty" xml:"item_detail,omitempty"`
	// 异常项指标详情
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
}
