package train

// TrainPassengerInfoDto 
type TrainPassengerInfoDto struct {

    // 实际出票价格（单位分）
    
    PayPrice   int64 `json:"pay_price,omitempty" xml:"pay_price,omitempty"`
    

    // 子订单号
    
    SubOrderId   int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
    

    // 改签手续费（单位分）
    
    HandFee   int64 `json:"hand_fee,omitempty" xml:"hand_fee,omitempty"`
    

    // 座位号
    
    SeatNum   string `json:"seat_num,omitempty" xml:"seat_num,omitempty"`
    

    // 真实坐席信息
    
    RealSeat   int64 `json:"real_seat,omitempty" xml:"real_seat,omitempty"`
    

    // 旅客类型
    
    PassengerType   int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
    

    // 乘客名称
    
    PassengerName   string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
    

    // 证件编号
    
    CertificateNum   string `json:"certificate_num,omitempty" xml:"certificate_num,omitempty"`
    

    // 证件类型
    
    CertificateType   int64 `json:"certificate_type,omitempty" xml:"certificate_type,omitempty"`
    

    // 儿童跟随，为成人的证件号
    
    ChdFollowAdtId   string `json:"chd_follow_adt_id,omitempty" xml:"chd_follow_adt_id,omitempty"`
    

}
