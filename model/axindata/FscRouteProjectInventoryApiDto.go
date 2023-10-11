package axindata

// FscRouteProjectInventoryApiDto 结构体
type FscRouteProjectInventoryApiDto struct {
	// 出团日期 格式yyyy-MM-dd
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 团期计划编号(供应商唯一团期标识)
	ProjectCode string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// 剩余库存数量（余位）
	InvCount int64 `json:"inv_count,omitempty" xml:"inv_count,omitempty"`
	// 已占库存数量
	OccupyCount int64 `json:"occupy_count,omitempty" xml:"occupy_count,omitempty"`
}
