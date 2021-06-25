package xhotel

// Topadshtlservicedatalist 
type Topadshtlservicedatalist struct {

    // 酒店id
    Hid   int64 `json:"hid,omitempty"`

    // 渠道名称
    Vendor   string `json:"vendor,omitempty"`

    // 电话
    Tel   string `json:"tel,omitempty"`

    // 有效订单数
    SalesCountInCycle   int64 `json:"sales_count_in_cycle,omitempty"`

    // 预定成功率
    BookingSuccRateStr   string `json:"booking_succ_rate_str,omitempty"`

    // 关闭订单数
    ClosedCountInCycle   int64 `json:"closed_count_in_cycle,omitempty"`

    // 查无订单数
    NoOrdCntInCycle   int64 `json:"no_ord_cnt_in_cycle,omitempty"`

    // 酒店名称
    Name   string `json:"name,omitempty"`

    // 统计时间
    ReportDate   string `json:"report_date,omitempty"`

    // 到店无房数量
    NoRoomCntInCycle   int64 `json:"no_room_cnt_in_cycle,omitempty"`

    // 酒店外部编码
    OuterId   string `json:"outer_id,omitempty"`

    // 是否查无订单 N:否  Y:是
    IsCallNoOrder   string `json:"is_call_no_order,omitempty"`

    // 是否特殊时段订单 N:否  Y:是
    IsSpecTimeOrder   string `json:"is_spec_time_order,omitempty"`

    // 入住时间
    CheckinDate   string `json:"checkin_date,omitempty"`

    // 是否到店无房订单  1:是
    IsNoRoomCompen   string `json:"is_no_room_compen,omitempty"`

    // 酒店名称
    HotelName   string `json:"hotel_name,omitempty"`

    // 确认时长
    ConfirmDuration   int64 `json:"confirm_duration,omitempty"`

    // 是否拒单 N:否  Y:是
    IsSellerDeny   string `json:"is_seller_deny,omitempty"`

    // 离店时间
    CheckoutDate   string `json:"checkout_date,omitempty"`

    // 是否卖家原因退款 N:否  Y:是
    IsSellerRefund   string `json:"is_seller_refund,omitempty"`

    // 供应商名称
    Supplier   string `json:"supplier,omitempty"`

    // 订单id
    Tid   int64 `json:"tid,omitempty"`

}
