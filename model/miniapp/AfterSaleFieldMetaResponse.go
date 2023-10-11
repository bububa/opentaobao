package miniapp

// AfterSaleFieldMetaResponse 结构体
type AfterSaleFieldMetaResponse struct {
	//
	Record []AfterSaleFieldMetaRecord `json:"record,omitempty" xml:"record>after_sale_field_meta_record,omitempty"`
	// 返回结果条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
