package qimen

// ItemInventory 
/* model for simplify = false
type ItemInventory struct {

    // 奇门仓储字段,C123,string(50),,
    
    ItemCode   string `json:"itemCode,omitempty"`
    

    // 奇门仓储字段,C123,string(50),,
    
    WarehouseCode   string `json:"warehouseCode,omitempty"`
    

    // 奇门仓储字段,C123,string(50),,
    
    ChannelCode   string `json:"channelCode,omitempty"`
    

    // 奇门仓储字段,C123,string(50),,
    
    Quantity   string `json:"quantity,omitempty"`
    

    // 奇门仓储字段,C123,string(50),,
    
    LockQuantity   string `json:"lockQuantity,omitempty"`
    

    // 奇门仓储字段
    
    OrderSourceCode   string `json:"orderSourceCode,omitempty"`
    

    // 奇门仓储字段
    
    SubSourceCode   string `json:"subSourceCode,omitempty"`
    

    // 奇门仓储字段
    
    ItemId   string `json:"itemId,omitempty"`
    

    // 响应结果:success|failure
    
    Flag   string `json:"flag,omitempty"`
    

    // 响应码
    
    Code   string `json:"code,omitempty"`
    

    // 响应信息
    
    Message   string `json:"message,omitempty"`
    

    // 奇门仓储字段
    
    CombItemId   string `json:"combItemId,omitempty"`
    

    // 奇门仓储字段
    
    ItemQuantity   string `json:"itemQuantity,omitempty"`
    

}
*/

// ItemInventory 
type ItemInventory struct {

    // 奇门仓储字段,C123,string(50),,
    ItemCode   string `json:"itemCode,omitempty"`

    // 奇门仓储字段,C123,string(50),,
    WarehouseCode   string `json:"warehouseCode,omitempty"`

    // 奇门仓储字段,C123,string(50),,
    ChannelCode   string `json:"channelCode,omitempty"`

    // 奇门仓储字段,C123,string(50),,
    Quantity   string `json:"quantity,omitempty"`

    // 奇门仓储字段,C123,string(50),,
    LockQuantity   string `json:"lockQuantity,omitempty"`

    // 奇门仓储字段
    OrderSourceCode   string `json:"orderSourceCode,omitempty"`

    // 奇门仓储字段
    SubSourceCode   string `json:"subSourceCode,omitempty"`

    // 奇门仓储字段
    ItemId   string `json:"itemId,omitempty"`

    // 响应结果:success|failure
    Flag   string `json:"flag,omitempty"`

    // 响应码
    Code   string `json:"code,omitempty"`

    // 响应信息
    Message   string `json:"message,omitempty"`

    // 奇门仓储字段
    CombItemId   string `json:"combItemId,omitempty"`

    // 奇门仓储字段
    ItemQuantity   string `json:"itemQuantity,omitempty"`

}
