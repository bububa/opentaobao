package flight

// RefundList 
/* model for simplify = false
type RefundList struct {

    // 乘客信息
    
    PassengerName   string `json:"passenger_name,omitempty"`
    

    // 退票费
    
    RefundFee   int64 `json:"refund_fee,omitempty"`
    

    // 票号
    
    Tickets  struct {
        String  []string `json:"string,omitempty"`
    } `json:"tickets,omitempty"`
    

    // 升舱退票费
    
    RefundUpgradeFee   int64 `json:"refund_upgrade_fee,omitempty"`
    

    // 改期退票费
    
    RefundModifyFee   int64 `json:"refund_modify_fee,omitempty"`
    

    // 退票航段
    
    RefundSegments  struct {
        RefundSegments  []RefundSegments `json:"refund_segments,omitempty"`
    } `json:"refund_segments,omitempty"`
    

    // 乘客类型
    
    PassengerType   int64 `json:"passenger_type,omitempty"`
    

    // 退票总额
    
    RefundAmount   int64 `json:"refund_amount,omitempty"`
    

    // 票面价
    
    TicketPrice   int64 `json:"ticket_price,omitempty"`
    

    // 税费
    
    Taxes  struct {
        Tax  []Tax `json:"tax,omitempty"`
    } `json:"taxes,omitempty"`
    

}
*/

// RefundList 
type RefundList struct {

    // 乘客信息
    PassengerName   string `json:"passenger_name,omitempty"`

    // 退票费
    RefundFee   int64 `json:"refund_fee,omitempty"`

    // 票号
    Tickets   []string `json:"tickets,omitempty"`

    // 升舱退票费
    RefundUpgradeFee   int64 `json:"refund_upgrade_fee,omitempty"`

    // 改期退票费
    RefundModifyFee   int64 `json:"refund_modify_fee,omitempty"`

    // 退票航段
    RefundSegments   []RefundSegments `json:"refund_segments,omitempty"`

    // 乘客类型
    PassengerType   int64 `json:"passenger_type,omitempty"`

    // 退票总额
    RefundAmount   int64 `json:"refund_amount,omitempty"`

    // 票面价
    TicketPrice   int64 `json:"ticket_price,omitempty"`

    // 税费
    Taxes   []Tax `json:"taxes,omitempty"`

}
