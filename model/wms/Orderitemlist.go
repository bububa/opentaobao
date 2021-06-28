package wms

// Orderitemlist 
/* model for simplify = false
type Orderitemlist struct {

    // 订单商品信息
    
    OrderItem  *struct {
        Orderitem  *Orderitem `json:"orderitem,omitempty"`
    } `json:"order_item,omitempty"`
    

}
*/

// Orderitemlist 
type Orderitemlist struct {

    // 订单商品信息
    OrderItem   *Orderitem `json:"order_item,omitempty"`

}
