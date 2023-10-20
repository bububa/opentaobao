package alihouse

// UpdateEcodeBuildingDto 结构体
type UpdateEcodeBuildingDto struct {
	// 外部私域楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部货-楼栋id（外部唯一识别码）
	OuterTid string `json:"outer_tid,omitempty" xml:"outer_tid,omitempty"`
	// 外部项目店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 楼栋E码
	ECode string `json:"e_code,omitempty" xml:"e_code,omitempty"`
}
