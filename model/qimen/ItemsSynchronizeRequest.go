package qimen

// ItemsSynchronizeRequest 
type ItemsSynchronizeRequest struct {

    // 操作类型(两种类型：add|update)
    ActionType   string `json:"actionType,omitempty"`

    // 仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
    WarehouseCode   string `json:"warehouseCode,omitempty"`

    // 货主编码
    OwnerCode   string `json:"ownerCode,omitempty"`

    // 同步的商品信息
    Items   []Item `json:"items,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
