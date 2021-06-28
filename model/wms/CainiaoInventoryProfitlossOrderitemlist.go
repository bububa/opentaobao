package wms

// CainiaoInventoryProfitlossOrderitemlist 
/* model for simplify = false
type CainiaoInventoryProfitlossOrderitemlist struct {

    // 损益详情
    
    OrderItem  *struct {
        CainiaoInventoryProfitlossOrderitem  *CainiaoInventoryProfitlossOrderitem `json:"cainiao_inventory_profitloss_orderitem,omitempty"`
    } `json:"order_item,omitempty"`
    

}
*/

// CainiaoInventoryProfitlossOrderitemlist 
type CainiaoInventoryProfitlossOrderitemlist struct {

    // 损益详情
    OrderItem   *CainiaoInventoryProfitlossOrderitem `json:"order_item,omitempty"`

}
