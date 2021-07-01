package damai

// FloorIdOpenParam 结构体
type FloorIdOpenParam struct {
	// 楼层id
	FloorId int64 `json:"floor_id,omitempty" xml:"floor_id,omitempty"`
	// 参数performId
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 参数systemId
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
}
