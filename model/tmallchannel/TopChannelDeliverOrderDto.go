package tmallchannel

// TopChannelDeliverOrderDto 
type TopChannelDeliverOrderDto struct {

    // 发货单单号
    
    MainDeliverOrderNo   int64 `json:"main_deliver_order_no,omitempty" xml:"main_deliver_order_no,omitempty"`
    

    // 发货单状态
    
    OrderStatus   int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
    

    // 审核状态
    
    AuditStatus   int64 `json:"audit_status,omitempty" xml:"audit_status,omitempty"`
    

    // 分销商Nick
    
    DistributorNick   string `json:"distributor_nick,omitempty" xml:"distributor_nick,omitempty"`
    

    // 渠道
    
    Channel   int64 `json:"channel,omitempty" xml:"channel,omitempty"`
    

    // 创建时间
    
    OrderCreateTime   string `json:"order_create_time,omitempty" xml:"order_create_time,omitempty"`
    

    // 最后更新时间
    
    OrderLastModifyTime   string `json:"order_last_modify_time,omitempty" xml:"order_last_modify_time,omitempty"`
    

    // 子发货单列表
    
    SubDeliverOrderList   []TopChannelSubDeliverOrderDto `json:"sub_deliver_order_list,omitempty" xml:"sub_deliver_order_list,omitempty"`
    

    // 物流单列表
    
    LogisticsOrderList   []TopChannelLogisticsOrderDto `json:"logistics_order_list,omitempty" xml:"logistics_order_list,omitempty"`
    

    // 解析详情
    
    Schema   string `json:"schema,omitempty" xml:"schema,omitempty"`
    

}
