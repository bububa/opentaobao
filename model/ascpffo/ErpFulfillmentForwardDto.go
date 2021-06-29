package ascpffo

// ErpFulfillmentForwardDTO 
type ErpFulfillmentForwardDTO struct {
    // 订单产生时间戳
    TradeCreateTime   int64 `json:"trade_create_time,omitempty" xml:"trade_create_time,omitempty"`
    // 发货金额
    PackagePaidFee   string `json:"package_paid_fee,omitempty" xml:"package_paid_fee,omitempty"`
    // 收货地区
    ReceiverCountry   string `json:"receiver_country,omitempty" xml:"receiver_country,omitempty"`
    // 收货人电话
    ReceiverPhone   string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
    // 收货人手机
    ReceiverMobile   string `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
    // 收货人姓名
    ReceiverName   string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
    // 买家昵称
    BuyerName   string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
    // 运单号
    TrackingNo   string `json:"tracking_no,omitempty" xml:"tracking_no,omitempty"`
    // 物流单号
    LbxNo   string `json:"lbx_no,omitempty" xml:"lbx_no,omitempty"`
    // 履约单号
    FulfillmentOrderNo   string `json:"fulfillment_order_no,omitempty" xml:"fulfillment_order_no,omitempty"`
    // 用户订单号
    TradeOrderNo   string `json:"trade_order_no,omitempty" xml:"trade_order_no,omitempty"`
    // 仓名称
    WarehouseName   string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
    // 出库时间戳
    OutBoundTime   int64 `json:"out_bound_time,omitempty" xml:"out_bound_time,omitempty"`
    // 下发到仓时间戳
    SendFulfillTime   int64 `json:"send_fulfill_time,omitempty" xml:"send_fulfill_time,omitempty"`
    // 订单状态
    OrderStatus   string `json:"order_status,omitempty" xml:"order_status,omitempty"`
    // 扩展字段
    ExtendFields   string `json:"extend_fields,omitempty" xml:"extend_fields,omitempty"`
}
