package flight

// PushAuxProductsRq 
type PushAuxProductsRq struct {

    // 接口身份标识用户名（渠道唯一标识）
    Cid   string `json:"cid,omitempty"`

    // 渠道id
    ChannelId   int64 `json:"channel_id,omitempty"`

    // 代理商ID
    AgentId   int64 `json:"agent_id,omitempty"`

    // 廉航辅营产品
    ProductItems   []AuxProductItemApiBean `json:"product_items,omitempty"`

}
