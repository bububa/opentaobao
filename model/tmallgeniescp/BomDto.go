package tmallgeniescp

// BomDto 结构体
type BomDto struct {
	// 二级物料-物料编码
	RawCode string `json:"raw_code,omitempty" xml:"raw_code,omitempty"`
	// sourceId
	SourceId string `json:"source_id,omitempty" xml:"source_id,omitempty"`
	// 成品-物料编码
	PrdId string `json:"prd_id,omitempty" xml:"prd_id,omitempty"`
	// 地点编码(成品供应商code)
	LocationCode string `json:"location_code,omitempty" xml:"location_code,omitempty"`
	// 扩展参数
	ExtendJson string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
	// 租户
	Tenant string `json:"tenant,omitempty" xml:"tenant,omitempty"`
	// 二级物料-行项目ID (10\20\30\40\50)
	SourceItemId string `json:"source_item_id,omitempty" xml:"source_item_id,omitempty"`
	// 二级物料-数量 (保留2位小数的数字)
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
}
