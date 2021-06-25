package wms

// CainiaoReturnBillOrderitem 
type CainiaoReturnBillOrderitem struct {

    // 商品ID
    OrderItemId   string `json:"order_item_id,omitempty"`

    // 商品ID
    ItemId   string `json:"item_id,omitempty"`

    // 商家编码
    ItemCode   string `json:"item_code,omitempty"`

    // 商品信息
    InventoryItemList   []CainiaoReturnBillInventoryitemlist `json:"inventory_item_list,omitempty"`

}
