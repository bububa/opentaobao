package ascpchannel

// Attribute 结构体
type Attribute struct {
	// 供应商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// FULLAMOUNT:全量模式；INCREMENT:增量模式
	InvOperateMode string `json:"inv_operate_mode,omitempty" xml:"inv_operate_mode,omitempty"`
}
