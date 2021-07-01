package wdk

// BomItemInfos 结构体
type BomItemInfos struct {
	// quantity
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// itemCode
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// cabinetCode
	CabinetCode string `json:"cabinet_code,omitempty" xml:"cabinet_code,omitempty"`
}
