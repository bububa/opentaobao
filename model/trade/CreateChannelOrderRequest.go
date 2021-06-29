package trade

// CreateChannelOrderRequest 
type CreateChannelOrderRequest struct {
    // 备注
    Memo   string `json:"memo,omitempty" xml:"memo,omitempty"`
    // 请求单号
    RequestNo   string `json:"request_no,omitempty" xml:"request_no,omitempty"`
    // 外部订单号
    OutOrderNo   string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
    // 子订单信息
    ItemList   []ChannelOrderItem `json:"item_list,omitempty" xml:"item_list>channel_order_item,omitempty"`
    // 自营实体标示
    BizType   int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
    // 物流信息
    ReceiverLogistics   *ReceiverLogistics `json:"receiver_logistics,omitempty" xml:"receiver_logistics,omitempty"`
    // sourceLbx
    SourceLbx   string `json:"source_lbx,omitempty" xml:"source_lbx,omitempty"`
    // 属性
    Properties   string `json:"properties,omitempty" xml:"properties,omitempty"`
    // 渠道
    Channel   int64 `json:"channel,omitempty" xml:"channel,omitempty"`
    // 选项
    Option   *ChannelOrderOption `json:"option,omitempty" xml:"option,omitempty"`
    // 交易类型（1——代销；2——经销）
    TradeType   int64 `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
}
