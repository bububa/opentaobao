package axindata

// FscPoiQueryRequest 结构体
type FscPoiQueryRequest struct {
	// 关键字
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 供应商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 分页大小（最大100）
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 是否国外景点
	Abroad bool `json:"abroad,omitempty" xml:"abroad,omitempty"`
}
