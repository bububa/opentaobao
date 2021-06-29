package axintrade

// AxinFundDTO 
type AxinFundDTO struct {
    // 有效资金单列表
    FundList   []AxinFundDto `json:"fund_list,omitempty" xml:"fund_list>axin_fund_dto,omitempty"`
    // 已支付总金额
    TotalPayedAmount   int64 `json:"total_payed_amount,omitempty" xml:"total_payed_amount,omitempty"`
}
