package wms

// Tmsorderlist 结构体
type Tmsorderlist struct {
	// 运单信息列表
	TmsOrder *Tmsorder `json:"tms_order,omitempty" xml:"tms_order,omitempty"`
}
