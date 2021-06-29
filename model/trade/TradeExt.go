package trade

// TradeExt 
type TradeExt struct {
    // enable前扩展标识位
    BeforeEnableFlag   int64 `json:"before_enable_flag,omitempty" xml:"before_enable_flag,omitempty"`
    // 关闭订单前扩展标识位
    BeforeCloseFlag   int64 `json:"before_close_flag,omitempty" xml:"before_close_flag,omitempty"`
    // 付款前扩展标识位
    BeforePayFlag   int64 `json:"before_pay_flag,omitempty" xml:"before_pay_flag,omitempty"`
    // 发货前扩展标识位
    BeforeShipFlag   int64 `json:"before_ship_flag,omitempty" xml:"before_ship_flag,omitempty"`
    // 确认收货前扩展标识位
    BeforeConfirmFlag   int64 `json:"before_confirm_flag,omitempty" xml:"before_confirm_flag,omitempty"`
    // 评价前扩展标识位
    BeforeRateFlag   int64 `json:"before_rate_flag,omitempty" xml:"before_rate_flag,omitempty"`
    // 退款前扩展标识位
    BeforeRefundFlag   int64 `json:"before_refund_flag,omitempty" xml:"before_refund_flag,omitempty"`
    // 修改前扩展标识位
    BeforeModifyFlag   int64 `json:"before_modify_flag,omitempty" xml:"before_modify_flag,omitempty"`
    // 第三方状态，第三方自由定义
    ThirdPartyStatus   int64 `json:"third_party_status,omitempty" xml:"third_party_status,omitempty"`
    // 第三方个性化数据
    ExtraData   string `json:"extra_data,omitempty" xml:"extra_data,omitempty"`
    // attributes标记
    ExtAttributes   string `json:"ext_attributes,omitempty" xml:"ext_attributes,omitempty"`
}
