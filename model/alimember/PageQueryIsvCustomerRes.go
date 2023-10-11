package alimember

// PageQueryIsvCustomerRes 结构体
type PageQueryIsvCustomerRes struct {
	// 会员数据
	Data []SaasCustomerInfo `json:"data,omitempty" xml:"data>saas_customer_info,omitempty"`
	// 总大小
	TotalSize int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
	// 当前第几页
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 是否还有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}
