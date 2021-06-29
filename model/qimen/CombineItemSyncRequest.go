package qimen

// CombineItemSyncRequest 
type CombineItemSyncRequest struct {
    // 组合商品的ERP编码
    ItemCode   string `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
    // 货主编码
    OwnerCode   string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
    // 仓库编码
    WarehouseCode   string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
    // 组合商品接口中的单商品信息
    Items   []Item `json:"items,omitempty" xml:"items>item,omitempty"`
    // 扩展属性
    ExtendProps   *TaobaoQimenCombineitemSynchronizeMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
