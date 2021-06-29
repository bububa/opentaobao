package wdk

// SkuSeriesEditRequest 
type SkuSeriesEditRequest struct {
    // 系列描述
    SeriesDesc   string `json:"series_desc,omitempty" xml:"series_desc,omitempty"`
    // 系列名称
    SeriesName   string `json:"series_name,omitempty" xml:"series_name,omitempty"`
    // 系列编码
    SeriesId   int64 `json:"series_id,omitempty" xml:"series_id,omitempty"`
}
