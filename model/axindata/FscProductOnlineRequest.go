package axindata

// FscProductOnlineRequest 结构体
type FscProductOnlineRequest struct {
	// 产品编号列表
	ProductCodeList []string `json:"product_code_list,omitempty" xml:"product_code_list>string,omitempty"`
	// 供应商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
}
