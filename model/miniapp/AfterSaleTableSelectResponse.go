package miniapp

// AfterSaleTableSelectResponse 结构体
type AfterSaleTableSelectResponse struct {
	// 查询结果
	JsonRecord string `json:"json_record,omitempty" xml:"json_record,omitempty"`
	// 查询结果条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
