package aecreatives

// TrafficImageSearchResultDto 结构体
type TrafficImageSearchResultDto struct {
	// 图搜结果
	Products []Product `json:"products,omitempty" xml:"products>product,omitempty"`
	// 总数
	TotalRecordCount int64 `json:"total_record_count,omitempty" xml:"total_record_count,omitempty"`
	// 图片识别的坐标
	Region *ProductImgRegionDto `json:"region,omitempty" xml:"region,omitempty"`
}
