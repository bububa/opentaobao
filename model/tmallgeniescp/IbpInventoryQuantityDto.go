package tmallgeniescp

// IbpInventoryQuantityDto 结构体
type IbpInventoryQuantityDto struct {
	// 扩展参数
	ExtendJson string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
	// 租户
	Tenant string `json:"tenant,omitempty" xml:"tenant,omitempty"`
	// 库存量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 地点code
	LocationCode string `json:"location_code,omitempty" xml:"location_code,omitempty"`
	// 物料code
	MaterielCode string `json:"materiel_code,omitempty" xml:"materiel_code,omitempty"`
}
