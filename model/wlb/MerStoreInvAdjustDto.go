package wlb

// MerStoreInvAdjustDto 
/* model for simplify = false
type MerStoreInvAdjustDto struct {

    // 库存类型
    
    InventoryType   int64 `json:"inventory_type,omitempty"`
    

    // 扩展属性
    
    Attribute   string `json:"attribute,omitempty"`
    

    // 数量
    
    Quantity   int64 `json:"quantity,omitempty"`
    

    // 外部操作唯一编码
    
    OutBizCode   string `json:"out_biz_code,omitempty"`
    

    // 仓库编码
    
    StoreCode   string `json:"store_code,omitempty"`
    

    // 货品id
    
    ScItemId   int64 `json:"sc_item_id,omitempty"`
    

}
*/

// MerStoreInvAdjustDto 
type MerStoreInvAdjustDto struct {

    // 库存类型
    InventoryType   int64 `json:"inventory_type,omitempty"`

    // 扩展属性
    Attribute   string `json:"attribute,omitempty"`

    // 数量
    Quantity   int64 `json:"quantity,omitempty"`

    // 外部操作唯一编码
    OutBizCode   string `json:"out_biz_code,omitempty"`

    // 仓库编码
    StoreCode   string `json:"store_code,omitempty"`

    // 货品id
    ScItemId   int64 `json:"sc_item_id,omitempty"`

}
