package wms

// CainiaoStockInBillOrderitem 
/* model for simplify = false
type CainiaoStockInBillOrderitem struct {

    // ERP订单明细ID
    
    OrderItemId   string `json:"order_item_id,omitempty"`
    

    // 商品ID
    
    ItemId   string `json:"item_id,omitempty"`
    

    // 商家编码
    
    ItemCode   string `json:"item_code,omitempty"`
    

    // 通知数量
    
    ApplyQty   int64 `json:"apply_qty,omitempty"`
    

    // 仓库收货商品信息
    
    InventoryItemList  struct {
        CainiaoStockInBillInventoryitemlist  []CainiaoStockInBillInventoryitemlist `json:"cainiao_stock_in_bill_inventoryitemlist,omitempty"`
    } `json:"inventory_item_list,omitempty"`
    

}
*/

// CainiaoStockInBillOrderitem 
type CainiaoStockInBillOrderitem struct {

    // ERP订单明细ID
    OrderItemId   string `json:"order_item_id,omitempty"`

    // 商品ID
    ItemId   string `json:"item_id,omitempty"`

    // 商家编码
    ItemCode   string `json:"item_code,omitempty"`

    // 通知数量
    ApplyQty   int64 `json:"apply_qty,omitempty"`

    // 仓库收货商品信息
    InventoryItemList   []CainiaoStockInBillInventoryitemlist `json:"inventory_item_list,omitempty"`

}
