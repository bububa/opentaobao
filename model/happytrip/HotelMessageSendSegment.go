package happytrip

// HotelMessageSendSegment 
type HotelMessageSendSegment struct {
    // 入住日期
    Checkin   string `json:"checkin,omitempty" xml:"checkin,omitempty"`
    // 离店日期
    Checkout   string `json:"checkout,omitempty" xml:"checkout,omitempty"`
    // 实际入住日期
    FeeBeginTime   string `json:"fee_begin_time,omitempty" xml:"fee_begin_time,omitempty"`
    // 实际离店日期
    FeeFinishTime   string `json:"fee_finish_time,omitempty" xml:"fee_finish_time,omitempty"`
    // 消息类型
    MessageType   string `json:"message_type,omitempty" xml:"message_type,omitempty"`
    // 外部订单号
    OutOrderId   string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
    // 消息结果
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
    // 房型
    RoomTypeName   string `json:"room_type_name,omitempty" xml:"room_type_name,omitempty"`
    // 供应商CODE
    SupplierCode   string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
    // 消息体参数容器
    Parameters   string `json:"parameters,omitempty" xml:"parameters,omitempty"`
}
