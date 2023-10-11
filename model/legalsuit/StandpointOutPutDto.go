package legalsuit

// StandpointOutPutDto 结构体
type StandpointOutPutDto struct {
	// 表格
	TableStr []string `json:"table_str,omitempty" xml:"table_str>string,omitempty"`
	// 标签
	StandpointLabels []string `json:"standpoint_labels,omitempty" xml:"standpoint_labels>string,omitempty"`
	// 拓展字段
	Options []Options `json:"options,omitempty" xml:"options>options,omitempty"`
	// 场景结构
	ScenesStruct string `json:"scenes_struct,omitempty" xml:"scenes_struct,omitempty"`
	// 标签
	Label string `json:"label,omitempty" xml:"label,omitempty"`
	// 创建日期
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 表格
	TableSchema string `json:"table_schema,omitempty" xml:"table_schema,omitempty"`
	// 口径
	DefenseCaliber string `json:"defense_caliber,omitempty" xml:"defense_caliber,omitempty"`
	// 描述
	StandpointDesc string `json:"standpoint_desc,omitempty" xml:"standpoint_desc,omitempty"`
	// 场景名称
	ScenesName string `json:"scenes_name,omitempty" xml:"scenes_name,omitempty"`
	// 创建人
	CreaterWorkerNo string `json:"creater_worker_no,omitempty" xml:"creater_worker_no,omitempty"`
	// 场景结构
	ScenesStrucct string `json:"scenes_strucct,omitempty" xml:"scenes_strucct,omitempty"`
	// 场景id
	ScenesId int64 `json:"scenes_id,omitempty" xml:"scenes_id,omitempty"`
	// 衍生口径数量
	DeriveCount int64 `json:"derive_count,omitempty" xml:"derive_count,omitempty"`
	// 引用数量
	ReferencedCount int64 `json:"referenced_count,omitempty" xml:"referenced_count,omitempty"`
	// 口径id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 是否收藏
	IsCollection bool `json:"is_collection,omitempty" xml:"is_collection,omitempty"`
}
