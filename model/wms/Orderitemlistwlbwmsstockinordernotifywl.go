package wms

// Orderitemlistwlbwmsstockinordernotifywl 
/* model for simplify = false
type Orderitemlistwlbwmsstockinordernotifywl struct {

    // 订单商品
    
    OrderItem  *struct {
        Orderitemwlbwmsstockinordernotifywl  *Orderitemwlbwmsstockinordernotifywl `json:"orderitemwlbwmsstockinordernotifywl,omitempty"`
    } `json:"order_item,omitempty"`
    

}
*/

// Orderitemlistwlbwmsstockinordernotifywl 
type Orderitemlistwlbwmsstockinordernotifywl struct {

    // 订单商品
    OrderItem   *Orderitemwlbwmsstockinordernotifywl `json:"order_item,omitempty"`

}
