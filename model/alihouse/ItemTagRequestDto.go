package alihouse

// ItemTagRequestDto 结构体
type ItemTagRequestDto struct {
	// 外部主键ID
	OuterIds []string `json:"outer_ids,omitempty" xml:"outer_ids>string,omitempty"`
	// 打到IC标，如果业务有db 比如楼盘会存db feature字段
	ItemTagCodes []string `json:"item_tag_codes,omitempty" xml:"item_tag_codes>string,omitempty"`
	// 外部门店名称
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 1-新房楼盘 2-小区
	SourceType int64 `json:"source_type,omitempty" xml:"source_type,omitempty"`
	// etc版本号
	EtcVersion int64 `json:"etc_version,omitempty" xml:"etc_version,omitempty"`
	// 默认0  1添加标签 -1去除标签
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
