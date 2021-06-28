package trade

// CreateChannelOrderRequest 
/* model for simplify = false
type CreateChannelOrderRequest struct {

    // 备注
    
    Memo   string `json:"memo,omitempty"`
    

    // 请求单号
    
    RequestNo   string `json:"request_no,omitempty"`
    

    // 外部订单号
    
    OutOrderNo   string `json:"out_order_no,omitempty"`
    

    // 子订单信息
    
    ItemList  struct {
        ChannelOrderItem  []ChannelOrderItem `json:"channel_order_item,omitempty"`
    } `json:"item_list,omitempty"`
    

    // 自营实体标示
    
    BizType   int64 `json:"biz_type,omitempty"`
    

    // 物流信息
    
    ReceiverLogistics  *struct {
        ReceiverLogistics  *ReceiverLogistics `json:"receiver_logistics,omitempty"`
    } `json:"receiver_logistics,omitempty"`
    

    // sourceLbx
    
    SourceLbx   string `json:"source_lbx,omitempty"`
    

    // 属性
    
    Properties   string `json:"properties,omitempty"`
    

    // 渠道
    
    Channel   int64 `json:"channel,omitempty"`
    

    // 选项
    
    Option  *struct {
        ChannelOrderOption  *ChannelOrderOption `json:"channel_order_option,omitempty"`
    } `json:"option,omitempty"`
    

    // 交易类型（1——代销；2——经销）
    
    TradeType   int64 `json:"trade_type,omitempty"`
    

}
*/

// CreateChannelOrderRequest 
type CreateChannelOrderRequest struct {

    // 备注
    Memo   string `json:"memo,omitempty"`

    // 请求单号
    RequestNo   string `json:"request_no,omitempty"`

    // 外部订单号
    OutOrderNo   string `json:"out_order_no,omitempty"`

    // 子订单信息
    ItemList   []ChannelOrderItem `json:"item_list,omitempty"`

    // 自营实体标示
    BizType   int64 `json:"biz_type,omitempty"`

    // 物流信息
    ReceiverLogistics   *ReceiverLogistics `json:"receiver_logistics,omitempty"`

    // sourceLbx
    SourceLbx   string `json:"source_lbx,omitempty"`

    // 属性
    Properties   string `json:"properties,omitempty"`

    // 渠道
    Channel   int64 `json:"channel,omitempty"`

    // 选项
    Option   *ChannelOrderOption `json:"option,omitempty"`

    // 交易类型（1——代销；2——经销）
    TradeType   int64 `json:"trade_type,omitempty"`

}
