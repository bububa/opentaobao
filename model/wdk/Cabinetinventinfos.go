package wdk

// Cabinetinventinfos 结构体
type Cabinetinventinfos struct {
	// 货位编码
	CabinetCode string `json:"cabinet_code,omitempty" xml:"cabinet_code,omitempty"`
	// 货位总量
	RealInvent string `json:"real_invent,omitempty" xml:"real_invent,omitempty"`
}
