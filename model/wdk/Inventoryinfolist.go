package wdk

// Inventoryinfolist 结构体
type Inventoryinfolist struct {
	// 库存单位
	StorageUnit string `json:"storage_unit,omitempty" xml:"storage_unit,omitempty"`
	// 实物总量
	RealInvent string `json:"real_invent,omitempty" xml:"real_invent,omitempty"`
	// 商品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 货位库存详情
	CabinetInventInfos []Cabinetinventinfos `json:"cabinet_invent_infos,omitempty" xml:"cabinet_invent_infos>cabinetinventinfos,omitempty"`
}
