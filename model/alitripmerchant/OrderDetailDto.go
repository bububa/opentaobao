package alitripmerchant

// OrderDetailDto 
type OrderDetailDto struct {
    // 酒店电话
    HotelPhone   string `json:"hotel_phone,omitempty" xml:"hotel_phone,omitempty"`
    // 入住人数
    PersonNum   int64 `json:"person_num,omitempty" xml:"person_num,omitempty"`
    // 早餐类别
    BreakfastType   string `json:"breakfast_type,omitempty" xml:"breakfast_type,omitempty"`
    // 房间图片
    RoomPicture   string `json:"room_picture,omitempty" xml:"room_picture,omitempty"`
    // 预定时间
    PaymentDate   string `json:"payment_date,omitempty" xml:"payment_date,omitempty"`
    // 订单编号
    OrderCode   string `json:"order_code,omitempty" xml:"order_code,omitempty"`
    // 纬度
    Lat   string `json:"lat,omitempty" xml:"lat,omitempty"`
    // 经度
    Lon   string `json:"lon,omitempty" xml:"lon,omitempty"`
    // 入住人邮箱
    ContactEmail   string `json:"contact_email,omitempty" xml:"contact_email,omitempty"`
    // 入住人手机
    ContactPhone   string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
    // 入住人姓
    ContactFirstName   string `json:"contact_first_name,omitempty" xml:"contact_first_name,omitempty"`
    // 入住人名
    ContactLastName   string `json:"contact_last_name,omitempty" xml:"contact_last_name,omitempty"`
    // 入住天数
    Days   int64 `json:"days,omitempty" xml:"days,omitempty"`
    // 离店时间
    CheckOutDate   string `json:"check_out_date,omitempty" xml:"check_out_date,omitempty"`
    // 入住时间
    CheckInDate   string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
    // 房间数量
    RoomNumber   int64 `json:"room_number,omitempty" xml:"room_number,omitempty"`
    // 价格名称
    RpName   string `json:"rp_name,omitempty" xml:"rp_name,omitempty"`
    // 最大入住人数
    MaxCheckInNumber   int64 `json:"max_check_in_number,omitempty" xml:"max_check_in_number,omitempty"`
    // 房间床型描述
    BedTypeDesc   string `json:"bed_type_desc,omitempty" xml:"bed_type_desc,omitempty"`
    // 房间面积
    RoomArea   string `json:"room_area,omitempty" xml:"room_area,omitempty"`
    // 房间名称
    RoomName   string `json:"room_name,omitempty" xml:"room_name,omitempty"`
    // 酒店名称
    HotelName   string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
    // 费用明细对象
    PriceDetailDto   *PriceDetailDto `json:"price_detail_dto,omitempty" xml:"price_detail_dto,omitempty"`
    // 货币
    Currency   string `json:"currency,omitempty" xml:"currency,omitempty"`
    // 总价格
    TotalPrice   string `json:"total_price,omitempty" xml:"total_price,omitempty"`
    // 支付类型
    PayType   int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
    // 取消描述
    CancelDec   string `json:"cancel_dec,omitempty" xml:"cancel_dec,omitempty"`
    // 取消政策规则
    CancelRule   int64 `json:"cancel_rule,omitempty" xml:"cancel_rule,omitempty"`
    // 支付剩余时间
    PayRemainTime   int64 `json:"pay_remain_time,omitempty" xml:"pay_remain_time,omitempty"`
    // 订单状态
    OrderStatus   string `json:"order_status,omitempty" xml:"order_status,omitempty"`
    // 酒店房型信息
    RoomDetailDto   *RoomDetailDto `json:"room_detail_dto,omitempty" xml:"room_detail_dto,omitempty"`
    // 卖家房型id
    OuterRoomId   string `json:"outer_room_id,omitempty" xml:"outer_room_id,omitempty"`
    // 租户酒店id
    HotelId   string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
    // 卖家价格ID
    OutRateId   int64 `json:"out_rate_id,omitempty" xml:"out_rate_id,omitempty"`
    // 酒店地址描述
    HotelAddress   string `json:"hotel_address,omitempty" xml:"hotel_address,omitempty"`
    // 退款手续费
    RefundCostAmount   string `json:"refund_cost_amount,omitempty" xml:"refund_cost_amount,omitempty"`
    // 成人总数
    AdultNumber   int64 `json:"adult_number,omitempty" xml:"adult_number,omitempty"`
    // 儿童总数
    ChildrenNumber   int64 `json:"children_number,omitempty" xml:"children_number,omitempty"`
    // 每个房间入住人信息
    GuestByRoomDtos   []GuestByRoomDto `json:"guest_by_room_dtos,omitempty" xml:"guest_by_room_dtos>guest_by_room_dto,omitempty"`
    // 订单状态描述
    OrderStatusDesc   string `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
    // 支付渠道
    PaymentChannel   string `json:"payment_channel,omitempty" xml:"payment_channel,omitempty"`
    // 酒店预订确认参数
    PmsCode   string `json:"pms_code,omitempty" xml:"pms_code,omitempty"`
}
