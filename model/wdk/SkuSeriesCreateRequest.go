package wdk

// SkuSeriesCreateRequest 结构体
type SkuSeriesCreateRequest struct {
	// 系列品描述
	SeriesDesc string `json:"series_desc,omitempty" xml:"series_desc,omitempty"`
	// 系列品名称
	SeriesName string `json:"series_name,omitempty" xml:"series_name,omitempty"`
	// 行业类型
	IndustryType string `json:"industry_type,omitempty" xml:"industry_type,omitempty"`
	// 类目id
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
}
