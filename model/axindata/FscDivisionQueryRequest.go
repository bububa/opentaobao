package axindata

// FscDivisionQueryRequest 结构体
type FscDivisionQueryRequest struct {
	// 行政区划类型（DOMAIN:境内，ABROAD:出境）
	DivisionType string `json:"division_type,omitempty" xml:"division_type,omitempty"`
	// 关键字
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 供应商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 国家ID
	CountryId int64 `json:"country_id,omitempty" xml:"country_id,omitempty"`
	// 行政区划级别（2-国家; 4-城市）
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
	// 分页大小（最大100）
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
}
