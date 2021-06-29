package xhoteloffline

// RoomSettleInfo 
type RoomSettleInfo struct {

    // 房间号
    
    RoomNo   string `json:"room_no,omitempty" xml:"room_no,omitempty"`
    

    // 房间费（大于零）
    
    RoomFee   int64 `json:"room_fee,omitempty" xml:"room_fee,omitempty"`
    

    // 房间杂费（不能为负数）
    
    RoomOtherFee   int64 `json:"room_other_fee,omitempty" xml:"room_other_fee,omitempty"`
    

    // 房间杂费明细 （当房间杂费存在时候，此值不能空。格式与原有杂费格式otherFeeDetail的格式相同）eg;{"洗衣费":5000,"水吧":10000,"优惠":1000}
    
    RoomOtherFeeDetail   string `json:"room_other_fee_detail,omitempty" xml:"room_other_fee_detail,omitempty"`
    

    // 房间check in 时间
    
    RoomCheckIn   string `json:"room_check_in,omitempty" xml:"room_check_in,omitempty"`
    

    // 房间check out时间
    
    RoomCheckOut   string `json:"room_check_out,omitempty" xml:"room_check_out,omitempty"`
    

    // 日历价格（每个房间的日历价格，多间房结账必填）
    
    DailyPriceInfo   string `json:"daily_price_info,omitempty" xml:"daily_price_info,omitempty"`
    

    // 房间状态；1:未入住（担保且需扣款）；2:取消成功（卖家取消；担保noshow且双方协商一致不扣款）；3:已入住
    
    RoomStatus   string `json:"room_status,omitempty" xml:"room_status,omitempty"`
    

}
