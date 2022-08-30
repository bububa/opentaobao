package wdk

// Inventoryinfolist 结构体
type Inventoryinfolist struct {
	// realInvent
	RealInvent string `json:"real_invent,omitempty" xml:"real_invent,omitempty"`
	// storageUnit
	StorageUnit string `json:"storage_unit,omitempty" xml:"storage_unit,omitempty"`
	// itemCode
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
}
