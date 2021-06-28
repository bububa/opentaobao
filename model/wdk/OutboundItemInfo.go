package wdk

// OutboundItemInfo 
/* model for simplify = false
type OutboundItemInfo struct {

    // 批发单号
    
    WholesaleOrderNo   string `json:"wholesale_order_no,omitempty"`
    

    // 已废弃，请使用containers.production_date
    
    ProductionDate   string `json:"production_date,omitempty"`
    

    // 是否完结
    
    OutboundCompleted   bool `json:"outbound_completed,omitempty"`
    

    // 出库数量(为正数或零)
    
    OutboundQuantity   string `json:"outbound_quantity,omitempty"`
    

    // 商品编码
    
    SkuCode   string `json:"sku_code,omitempty"`
    

    // 外部单号，如采购单号
    
    ExternalOrderNo   string `json:"external_order_no,omitempty"`
    

    // 容器信息
    
    Containers  struct {
        ContainerDo  []ContainerDo `json:"container_do,omitempty"`
    } `json:"containers,omitempty"`
    

}
*/

// OutboundItemInfo 
type OutboundItemInfo struct {

    // 批发单号
    WholesaleOrderNo   string `json:"wholesale_order_no,omitempty"`

    // 已废弃，请使用containers.production_date
    ProductionDate   string `json:"production_date,omitempty"`

    // 是否完结
    OutboundCompleted   bool `json:"outbound_completed,omitempty"`

    // 出库数量(为正数或零)
    OutboundQuantity   string `json:"outbound_quantity,omitempty"`

    // 商品编码
    SkuCode   string `json:"sku_code,omitempty"`

    // 外部单号，如采购单号
    ExternalOrderNo   string `json:"external_order_no,omitempty"`

    // 容器信息
    Containers   []ContainerDo `json:"containers,omitempty"`

}
