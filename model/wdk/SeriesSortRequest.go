package wdk

// SeriesSortRequest 
type SeriesSortRequest struct {
    // 系列编码
    SeriesId   int64 `json:"series_id,omitempty" xml:"series_id,omitempty"`
    // 有序行业属性对：行业属性key，属性值
    IndustryAttrList   string `json:"industry_attr_list,omitempty" xml:"industry_attr_list,omitempty"`
    // 行业类型
    IndustryType   string `json:"industry_type,omitempty" xml:"industry_type,omitempty"`
}
