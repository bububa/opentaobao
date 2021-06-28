package wms

// CainiaoStockInBillInventoryitemlist 
/* model for simplify = false
type CainiaoStockInBillInventoryitemlist struct {

    // 仓库收货商品信息
    
    InventoryItem  *struct {
        CainiaoStockInBillInventoryitem  *CainiaoStockInBillInventoryitem `json:"cainiao_stock_in_bill_inventoryitem,omitempty"`
    } `json:"inventory_item,omitempty"`
    

}
*/

// CainiaoStockInBillInventoryitemlist 
type CainiaoStockInBillInventoryitemlist struct {

    // 仓库收货商品信息
    InventoryItem   *CainiaoStockInBillInventoryitem `json:"inventory_item,omitempty"`

}
