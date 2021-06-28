package flight

// RefundPassengerFeeParam 
type RefundPassengerFeeParam struct {

    // 机票已使用部分总价(单位:分)
    
    AlreadyUsedTotalPirce   int64 `json:"already_used_total_pirce,omitempty" xml:"already_used_total_pirce,omitempty"`
    

    // 机票不可退税费(单位:分)
    
    NonRefundableTaxPrice   int64 `json:"non_refundable_tax_price,omitempty" xml:"non_refundable_tax_price,omitempty"`
    

    // 机票不可退票价(单位:分)
    
    NonRefundableTicketPrice   int64 `json:"non_refundable_ticket_price,omitempty" xml:"non_refundable_ticket_price,omitempty"`
    

    // 机票不可退改签服务费(单位:分)
    
    NonRefundableTotalChangeServiceFee   int64 `json:"non_refundable_total_change_service_fee,omitempty" xml:"non_refundable_total_change_service_fee,omitempty"`
    

    // 机票不可退改签升舱费(单位:分)
    
    NonRefundableTotalChangeUpgradeFee   int64 `json:"non_refundable_total_change_upgrade_fee,omitempty" xml:"non_refundable_total_change_upgrade_fee,omitempty"`
    

    // 乘机人姓名
    
    PassengerName   string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
    

}
