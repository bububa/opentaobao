package wdk

// SkuSeriesCreateRequest 
/* model for simplify = false
type SkuSeriesCreateRequest struct {

    // 系列品描述
    
    SeriesDesc   string `json:"series_desc,omitempty"`
    

    // 系列品名称
    
    SeriesName   string `json:"series_name,omitempty"`
    

    // 行业类型
    
    IndustryType   string `json:"industry_type,omitempty"`
    

    // 类目id
    
    CategoryId   int64 `json:"category_id,omitempty"`
    

}
*/

// SkuSeriesCreateRequest 
type SkuSeriesCreateRequest struct {

    // 系列品描述
    SeriesDesc   string `json:"series_desc,omitempty"`

    // 系列品名称
    SeriesName   string `json:"series_name,omitempty"`

    // 行业类型
    IndustryType   string `json:"industry_type,omitempty"`

    // 类目id
    CategoryId   int64 `json:"category_id,omitempty"`

}
