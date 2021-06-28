package wms

// CainiaoStockOutBillOrderitem 
/* model for simplify = false
type CainiaoStockOutBillOrderitem struct {

    // order_item_id
    
    OrderItemId   string `json:"order_item_id,omitempty"`
    

    // 商品ID
    
    ItemId   string `json:"item_id,omitempty"`
    

    // 商家编码
    
    ItemCode   string `json:"item_code,omitempty"`
    

    // 通知数量
    
    ApplyQty   int64 `json:"apply_qty,omitempty"`
    

    // 商品属性列表
    
    InventoryItemList  struct {
        CainiaoStockOutBillInventoryitemlist  []CainiaoStockOutBillInventoryitemlist `json:"cainiao_stock_out_bill_inventoryitemlist,omitempty"`
    } `json:"inventory_item_list,omitempty"`
    

}
*/

// CainiaoStockOutBillOrderitem 
type CainiaoStockOutBillOrderitem struct {

    // order_item_id
    OrderItemId   string `json:"order_item_id,omitempty"`

    // 商品ID
    ItemId   string `json:"item_id,omitempty"`

    // 商家编码
    ItemCode   string `json:"item_code,omitempty"`

    // 通知数量
    ApplyQty   int64 `json:"apply_qty,omitempty"`

    // 商品属性列表
    InventoryItemList   []CainiaoStockOutBillInventoryitemlist `json:"inventory_item_list,omitempty"`

}
