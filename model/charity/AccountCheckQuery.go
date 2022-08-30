package charity

// AccountCheckQuery 结构体
type AccountCheckQuery struct {
	// 账单列表
	BillList []BillDto `json:"bill_list,omitempty" xml:"bill_list>bill_dto,omitempty"`
	// 公益组织id
	InvoiceOrgId string `json:"invoice_org_id,omitempty" xml:"invoice_org_id,omitempty"`
	// 开票申请id
	InvoiceId string `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// 商家id
	MerchantId int64 `json:"merchant_id,omitempty" xml:"merchant_id,omitempty"`
}
