package qimen

// CombineItemSyncRequest 
type CombineItemSyncRequest struct {

    // 组合商品的ERP编码
    ItemCode   string `json:"itemCode,omitempty"`

    // 货主编码
    OwnerCode   string `json:"ownerCode,omitempty"`

    // 仓库编码
    WarehouseCode   string `json:"warehouseCode,omitempty"`

    // 组合商品接口中的单商品信息
    Items   []Item `json:"items,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
