package legalcase

// Page 结构体
type Page struct {
	// 返回列表
	Datas []StandpointSearchDto `json:"datas,omitempty" xml:"datas>standpoint_search_dto,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
