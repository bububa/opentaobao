package miniapp

// AfterSaleGetWorkTableListResponse 结构体
type AfterSaleGetWorkTableListResponse struct {
	// 调用结果
	Record []AfterSaleGetWorkTableListRecord `json:"record,omitempty" xml:"record>after_sale_get_work_table_list_record,omitempty"`
	// 返回结果数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
