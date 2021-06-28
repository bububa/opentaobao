package trade

// FastBuyPosQueryRequest 
/* model for simplify = false
type FastBuyPosQueryRequest struct {

    // pos机id
    
    MachineId   string `json:"machine_id,omitempty"`
    

    // 外部订单id
    
    OutOrderId   string `json:"out_order_id,omitempty"`
    

    // 经营店id
    
    StoreId   string `json:"store_id,omitempty"`
    

    // 外部门店编码
    
    OutShopCode   string `json:"out_shop_code,omitempty"`
    

}
*/

// FastBuyPosQueryRequest 
type FastBuyPosQueryRequest struct {

    // pos机id
    MachineId   string `json:"machine_id,omitempty"`

    // 外部订单id
    OutOrderId   string `json:"out_order_id,omitempty"`

    // 经营店id
    StoreId   string `json:"store_id,omitempty"`

    // 外部门店编码
    OutShopCode   string `json:"out_shop_code,omitempty"`

}
