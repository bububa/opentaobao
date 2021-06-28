package flight

// ModifyItemDTO 
/* model for simplify = false
type ModifyItemDTO struct {

    // 乘客姓名
    
    PassengerName   string `json:"passenger_name,omitempty"`
    

    // 票号
    
    Tickets  struct {
        String  []string `json:"string,omitempty"`
    } `json:"tickets,omitempty"`
    

    // pnr
    
    Pnr   string `json:"pnr,omitempty"`
    

    // 改签后航段
    
    AfterChangeSegments  struct {
        ModifySegmentDTO  []ModifySegmentDTO `json:"modify_segment_dto,omitempty"`
    } `json:"after_change_segments,omitempty"`
    

    // 改签费用
    
    ModifyFee   int64 `json:"modify_fee,omitempty"`
    

    // 升舱费用
    
    UpgradeFee   int64 `json:"upgrade_fee,omitempty"`
    

    // 改前航段
    
    BeforeChangeSegments  struct {
        ModifyBeforeSegmentDTO  []ModifyBeforeSegmentDTO `json:"modify_before_segment_dto,omitempty"`
    } `json:"before_change_segments,omitempty"`
    

    // 票号
    
    TicketNos  struct {
        String  []string `json:"string,omitempty"`
    } `json:"ticket_nos,omitempty"`
    

}
*/

// ModifyItemDTO 
type ModifyItemDTO struct {

    // 乘客姓名
    PassengerName   string `json:"passenger_name,omitempty"`

    // 票号
    Tickets   []string `json:"tickets,omitempty"`

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
    TicketNos   []string `json:"ticket_nos,omitempty"`

}
