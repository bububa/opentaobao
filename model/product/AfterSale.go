package product

// AfterSale 结构体
type AfterSale struct {
	// 名称
	AfterSaleName string `json:"after_sale_name,omitempty" xml:"after_sale_name,omitempty"`
	// tfs地址
	AfterSalePath string `json:"after_sale_path,omitempty" xml:"after_sale_path,omitempty"`
	// id
	AfterSaleId int64 `json:"after_sale_id,omitempty" xml:"after_sale_id,omitempty"`
}
