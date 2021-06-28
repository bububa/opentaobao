package qimen

// OrderLines 
/* model for simplify = false
type OrderLines struct {

    // 订单详情
    
    OrderLine  *struct {
        OrderLine  *OrderLine `json:"order_line,omitempty"`
    } `json:"orderLine,omitempty"`
    

}
*/

// OrderLines 
type OrderLines struct {

    // 订单详情
    OrderLine   *OrderLine `json:"orderLine,omitempty"`

}
