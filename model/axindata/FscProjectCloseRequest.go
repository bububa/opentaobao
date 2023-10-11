package axindata

// FscProjectCloseRequest 结构体
type FscProjectCloseRequest struct {
	// 团期编号列表
	ProjectCodeList []string `json:"project_code_list,omitempty" xml:"project_code_list>string,omitempty"`
	// 产品编号
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 行程编号
	JourneyCode string `json:"journey_code,omitempty" xml:"journey_code,omitempty"`
	// 供应商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
}
