package wlbimports

// LocOrder 
type LocOrder struct {

    // 运费
    ShippingFee   int64 `json:"shipping_fee,omitempty"`

    // 物流运单号
    TrackingNo   string `json:"tracking_no,omitempty"`

    // 交易订单号
    TradeId   int64 `json:"trade_id,omitempty"`

    // 物流订单号
    OrderCode   string `json:"order_code,omitempty"`

    // 物流承运商
    Carrier   string `json:"carrier,omitempty"`

    // 物流订单状态编码
    StatusCode   string `json:"status_code,omitempty"`

    // 关税
    CustomsFee   int64 `json:"customs_fee,omitempty"`

    // 重量
    Weight   int64 `json:"weight,omitempty"`

    // 重量单位
    WeightUnit   string `json:"weight_unit,omitempty"`

    // 费用币种
    Currency   string `json:"currency,omitempty"`

    // 订单状态中文描述
    StatusCodeDesc   string `json:"status_code_desc,omitempty"`

}
