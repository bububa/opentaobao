package nrt

// PageData 结构体
type PageData struct {
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 数据列表
	DataList []CouponTemplateDto `json:"data_list,omitempty" xml:"data_list>coupon_template_dto,omitempty"`
	// 1
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
