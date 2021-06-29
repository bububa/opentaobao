package maitix

// MoaTicketInfo 
type MoaTicketInfo struct {

    // 套票组ID，表示某几个座位是一套.机选可不填。h5选座的话就传third_combine_joint_ticket_id的值
    
    CombineId   int64 `json:"combine_id,omitempty" xml:"combine_id,omitempty"`
    

    // 外部子订单号，分销方子订单号，可不填
    
    ExternalSubOrderNo   string `json:"external_sub_order_no,omitempty" xml:"external_sub_order_no,omitempty"`
    

    // 票的实际持有人证件号，一票一证必填
    
    RealTicketOwnerIdCardNo   string `json:"real_ticket_owner_id_card_no,omitempty" xml:"real_ticket_owner_id_card_no,omitempty"`
    

    // 证件类型 1身份证 2护照 3港澳通行证 4台胞证 5士兵/军官，一票一证必填
    
    RealTicketOwnerIdCardType   int64 `json:"real_ticket_owner_id_card_type,omitempty" xml:"real_ticket_owner_id_card_type,omitempty"`
    

    // 票的实际持有人姓名，一票一证必填
    
    RealTicketOwnerName   string `json:"real_ticket_owner_name,omitempty" xml:"real_ticket_owner_name,omitempty"`
    

    // 票的实际持有人电话，一票一证必填
    
    RealTicketOwnerPhone   string `json:"real_ticket_owner_phone,omitempty" xml:"real_ticket_owner_phone,omitempty"`
    

    // 票的实际持有人电话国家代码，一票一证必填
    
    RealTicketOwnerPhoneCountryCode   string `json:"real_ticket_owner_phone_country_code,omitempty" xml:"real_ticket_owner_phone_country_code,omitempty"`
    

    // 座位ID，有座选座项目必填。
    
    SeatId   int64 `json:"seat_id,omitempty" xml:"seat_id,omitempty"`
    

    // 票品ID,如果是套票就是套票的票品id.有的地方也叫price_id。必填
    
    TicketItemId   int64 `json:"ticket_item_id,omitempty" xml:"ticket_item_id,omitempty"`
    

}
