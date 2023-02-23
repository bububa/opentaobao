package alihouse

// ProjectSalesTimePrePermitDto 结构体
type ProjectSalesTimePrePermitDto struct {
	// 预售证外部id
	PreSalePermitId string `json:"pre_sale_permit_id,omitempty" xml:"pre_sale_permit_id,omitempty"`
	// 预售时间
	PreSalePermitTime string `json:"pre_sale_permit_time,omitempty" xml:"pre_sale_permit_time,omitempty"`
	// 预售证编号
	PreSalePermitNo string `json:"pre_sale_permit_no,omitempty" xml:"pre_sale_permit_no,omitempty"`
}
