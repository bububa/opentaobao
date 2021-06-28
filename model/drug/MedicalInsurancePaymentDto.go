package drug

// MedicalInsurancePaymentDto 
/* model for simplify = false
type MedicalInsurancePaymentDto struct {

    // 医保支付金额信息
    
    FundBillList  struct {
        FundBillDto  []FundBillDto `json:"fund_bill_dto,omitempty"`
    } `json:"fund_bill_list,omitempty"`
    

    // 医保局返回的医保支付内容
    
    IndustrySepcDetail   string `json:"industry_sepc_detail,omitempty"`
    

}
*/

// MedicalInsurancePaymentDto 
type MedicalInsurancePaymentDto struct {

    // 医保支付金额信息
    FundBillList   []FundBillDto `json:"fund_bill_list,omitempty"`

    // 医保局返回的医保支付内容
    IndustrySepcDetail   string `json:"industry_sepc_detail,omitempty"`

}
