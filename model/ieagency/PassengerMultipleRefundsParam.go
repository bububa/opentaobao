package ieagency

// PassengerMultipleRefundsParam 
type PassengerMultipleRefundsParam struct {

    // 补退金额(单位:分)
    
    RefundMoney   int64 `json:"refund_money,omitempty" xml:"refund_money,omitempty"`
    

    // 乘机人姓名
    
    PassengerName   string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
    

}
