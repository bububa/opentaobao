package txcs

// StatementBillConfirmDto 
type StatementBillConfirmDto struct {

    // 结算公司编码
    
    SettlementCompanyCode   string `json:"settlement_company_code,omitempty" xml:"settlement_company_code,omitempty"`
    

    // 幂等ID
    
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    

    // 操作人ID
    
    OperatorId   string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
    

    // 操作人名称
    
    OperatorName   string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
    

    // 账单code
    
    StatementBillCodes   []string `json:"statement_bill_codes,omitempty" xml:"statement_bill_codes>string,omitempty"`
    

}
