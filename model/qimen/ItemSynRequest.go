package qimen

// ItemSynRequest 
/* model for simplify = false
type ItemSynRequest struct {

    // 操作类型(两种类型：add|update)
    
    ActionType   string `json:"actionType,omitempty"`
    

    // 仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
    
    WarehouseCode   string `json:"warehouseCode,omitempty"`
    

    // 货主编码
    
    OwnerCode   string `json:"ownerCode,omitempty"`
    

    // 供应商编码
    
    SupplierCode   string `json:"supplierCode,omitempty"`
    

    // 供应商名称
    
    SupplierName   string `json:"supplierName,omitempty"`
    

    // 商品信息
    
    Item  *struct {
        Item  *Item `json:"item,omitempty"`
    } `json:"item,omitempty"`
    

    // 扩展属性
    
    ExtendProps  *struct {
        Map  *Map `json:"map,omitempty"`
    } `json:"extendProps,omitempty"`
    

}
*/

// ItemSynRequest 
type ItemSynRequest struct {

    // 操作类型(两种类型：add|update)
    ActionType   string `json:"actionType,omitempty"`

    // 仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
    WarehouseCode   string `json:"warehouseCode,omitempty"`

    // 货主编码
    OwnerCode   string `json:"ownerCode,omitempty"`

    // 供应商编码
    SupplierCode   string `json:"supplierCode,omitempty"`

    // 供应商名称
    SupplierName   string `json:"supplierName,omitempty"`

    // 商品信息
    Item   *Item `json:"item,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
