package wdk

// CouponStatisticsParamDo 结构体
type CouponStatisticsParamDo struct {
	// 页码，即当前第几页
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 每页记录数，不能超过200
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 日期，格式为yyyymmdd
	StatisticsDate string `json:"statistics_date,omitempty" xml:"statistics_date,omitempty"`
	// 品牌名称数组
	BrandNames []string `json:"brand_names,omitempty" xml:"brand_names>string,omitempty"`
}
