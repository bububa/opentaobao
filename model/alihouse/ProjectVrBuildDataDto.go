package alihouse

// ProjectVrBuildDataDto 结构体
type ProjectVrBuildDataDto struct {
	// 外部房源ID
	OuterHouseId string `json:"outer_house_id,omitempty" xml:"outer_house_id,omitempty"`
	// 外部小区ID
	OuterProjectId string `json:"outer_project_id,omitempty" xml:"outer_project_id,omitempty"`
	// 提取码
	ExtractedCode string `json:"extracted_code,omitempty" xml:"extracted_code,omitempty"`
	// 外部户型ID
	OuterLayoutId string `json:"outer_layout_id,omitempty" xml:"outer_layout_id,omitempty"`
	// 外部货户型ID
	OuterLayoutTid string `json:"outer_layout_tid,omitempty" xml:"outer_layout_tid,omitempty"`
	// 是否删除 1-是 0-否
	IsDeleted int64 `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
	// 类型:新房-1，二手房-2
	SourceType int64 `json:"source_type,omitempty" xml:"source_type,omitempty"`
}
