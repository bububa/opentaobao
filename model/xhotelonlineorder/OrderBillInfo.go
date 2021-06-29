package xhotelonlineorder

// OrderBillInfo 
type OrderBillInfo struct {
    // 未结账房费
    NoCheckoutPrice   int64 `json:"no_checkout_price,omitempty" xml:"no_checkout_price,omitempty"`
    // 已结账房费
    CheckoutPrice   int64 `json:"checkout_price,omitempty" xml:"checkout_price,omitempty"`
    // 已结账总费用
    CheckoutTotalFee   int64 `json:"checkout_total_fee,omitempty" xml:"checkout_total_fee,omitempty"`
    // 未结账总费用
    NoCheckoutTotalFee   int64 `json:"no_checkout_total_fee,omitempty" xml:"no_checkout_total_fee,omitempty"`
    // 离店日期
    CheckOutDate   string `json:"check_out_date,omitempty" xml:"check_out_date,omitempty"`
    // 入住日期
    CheckInDate   string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
    // 房间号
    RoomNo   string `json:"room_no,omitempty" xml:"room_no,omitempty"`
    // 完整的身份证号
    IdNumber   string `json:"id_number,omitempty" xml:"id_number,omitempty"`
    // 客人姓名
    GuestName   string `json:"guest_name,omitempty" xml:"guest_name,omitempty"`
    // 外部酒店代码
    HotelCode   string `json:"hotel_code,omitempty" xml:"hotel_code,omitempty"`
    // 外部订单号
    OutOrderId   string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
    // 淘宝订单号
    Tid   int64 `json:"tid,omitempty" xml:"tid,omitempty"`
    // 请求id (同入参)
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 未结账杂费
    NoCheckoutOtherFee   int64 `json:"no_checkout_other_fee,omitempty" xml:"no_checkout_other_fee,omitempty"`
    // 已结账杂费
    CheckoutOtherFee   int64 `json:"checkout_other_fee,omitempty" xml:"checkout_other_fee,omitempty"`
    // 备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    // 杂费明细列表
    OtherFeeDetail   *OtherFeeDetail `json:"other_fee_detail,omitempty" xml:"other_fee_detail,omitempty"`
    // 每日房费类表
    DailyRoomFee   *DailyRoomFee `json:"daily_room_fee,omitempty" xml:"daily_room_fee,omitempty"`
}
