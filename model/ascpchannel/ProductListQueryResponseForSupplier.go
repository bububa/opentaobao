package ascpchannel

// ProductListQueryResponseForSupplier 结构体
type ProductListQueryResponseForSupplier struct {
	// 产品列表
	ProductList []ProductList `json:"product_list,omitempty" xml:"product_list>product_list,omitempty"`
	// 总数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
