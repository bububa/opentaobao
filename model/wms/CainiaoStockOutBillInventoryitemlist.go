package wms

// CainiaoStockOutBillInventoryitemlist 
/* model for simplify = false
type CainiaoStockOutBillInventoryitemlist struct {

    // 商品属性
    
    InventoryItem  *struct {
        CainiaoStockOutBillInventoryitem  *CainiaoStockOutBillInventoryitem `json:"cainiao_stock_out_bill_inventoryitem,omitempty"`
    } `json:"inventory_item,omitempty"`
    

}
*/

// CainiaoStockOutBillInventoryitemlist 
type CainiaoStockOutBillInventoryitemlist struct {

    // 商品属性
    InventoryItem   *CainiaoStockOutBillInventoryitem `json:"inventory_item,omitempty"`

}
