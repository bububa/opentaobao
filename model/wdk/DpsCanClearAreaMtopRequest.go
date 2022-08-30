package wdk

// DpsCanClearAreaMtopRequest 结构体
type DpsCanClearAreaMtopRequest struct {
	// 波次号
	WaveCode string `json:"wave_code,omitempty" xml:"wave_code,omitempty"`
	// 仓code
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
}
