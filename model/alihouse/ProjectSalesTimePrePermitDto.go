package alihouse

// ProjectSalesTimePrePermitDto 结构体
type ProjectSalesTimePrePermitDto struct {
	// 预售证编号
	PreSalePermitId string `json:"pre_sale_permit_id,omitempty" xml:"pre_sale_permit_id,omitempty"`
	// 预售时间
	PreSalePermitTime string `json:"pre_sale_permit_time,omitempty" xml:"pre_sale_permit_time,omitempty"`
}
