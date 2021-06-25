package flight

// ModifyItemDTO 
type ModifyItemDTO struct {

    // 乘客姓名
    PassengerName   string `json:"passenger_name,omitempty"`

    // 票号
    Tickets   []String `json:"tickets,omitempty"`

    // pnr
    Pnr   string `json:"pnr,omitempty"`

    // 改签后航段
    AfterChangeSegments   []ModifySegmentDTO `json:"after_change_segments,omitempty"`

    // 改签费用
    ModifyFee   int64 `json:"modify_fee,omitempty"`

    // 升舱费用
    UpgradeFee   int64 `json:"upgrade_fee,omitempty"`

    // 改前航段
    BeforeChangeSegments   []ModifyBeforeSegmentDTO `json:"before_change_segments,omitempty"`

    // 票号
    TicketNos   []String `json:"ticket_nos,omitempty"`

}
