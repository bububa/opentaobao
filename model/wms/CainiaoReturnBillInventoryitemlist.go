package wms

// CainiaoReturnBillInventoryitemlist 
/* model for simplify = false
type CainiaoReturnBillInventoryitemlist struct {

    // 订单详情
    
    InventoryItem  *struct {
        CainiaoReturnBillInventoryitem  *CainiaoReturnBillInventoryitem `json:"cainiao_return_bill_inventoryitem,omitempty"`
    } `json:"inventory_item,omitempty"`
    

}
*/

// CainiaoReturnBillInventoryitemlist 
type CainiaoReturnBillInventoryitemlist struct {

    // 订单详情
    InventoryItem   *CainiaoReturnBillInventoryitem `json:"inventory_item,omitempty"`

}
