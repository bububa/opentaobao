package alihouse

// UpdateProjectLayoutEcodeDto 结构体
type UpdateProjectLayoutEcodeDto struct {
	// 楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部货ID
	OuterTid string `json:"outer_tid,omitempty" xml:"outer_tid,omitempty"`
	// 户型外部唯一ID
	OuterLayoutId string `json:"outer_layout_id,omitempty" xml:"outer_layout_id,omitempty"`
	// 外部项目店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// eCode
	ECode string `json:"e_code,omitempty" xml:"e_code,omitempty"`
	// 数据源类型 （1-新房 2-二手房）
	SourceType int64 `json:"source_type,omitempty" xml:"source_type,omitempty"`
}
