package wms

// Orderitemlistwlbwmsstockoutordernotify 
/* model for simplify = false
type Orderitemlistwlbwmsstockoutordernotify struct {

    // 订单商品信息
    
    OrderItem  *struct {
        Orderitemwlbwmsstockoutordernotify  *Orderitemwlbwmsstockoutordernotify `json:"orderitemwlbwmsstockoutordernotify,omitempty"`
    } `json:"order_item,omitempty"`
    

}
*/

// Orderitemlistwlbwmsstockoutordernotify 
type Orderitemlistwlbwmsstockoutordernotify struct {

    // 订单商品信息
    OrderItem   *Orderitemwlbwmsstockoutordernotify `json:"order_item,omitempty"`

}
