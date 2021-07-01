package qimen

// TaobaoQimenInventoryreserveCancelRequest 
type TaobaoQimenInventoryreserveCancelRequest struct {
    // 奇门仓储字段
    OwnerCode   string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
    // 奇门仓储字段
    OrderCode   string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
    // 奇门仓储字段
    OrderSource   string `json:"orderSource,omitempty" xml:"orderSource,omitempty"`
    // 奇门仓储字段
    ItemInventories   []ItemInventory `json:"itemInventories,omitempty" xml:"itemInventories>item_inventory,omitempty"`
    // 扩展属性
    ExtendProps   *TaobaoQimenInventoryreserveCancelMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
