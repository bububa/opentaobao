package drug

// MedicalInsurancePaymentDTO 
type MedicalInsurancePaymentDTO struct {
    // 医保支付金额信息
    FundBillList   []FundBillDTO `json:"fund_bill_list,omitempty" xml:"fund_bill_list>fund_bill_dto,omitempty"`
    // 医保局返回的医保支付内容
    IndustrySepcDetail   string `json:"industry_sepc_detail,omitempty" xml:"industry_sepc_detail,omitempty"`
}
