package wms

// Orderitemlistwlbwmsconsignordernotify 
/* model for simplify = false
type Orderitemlistwlbwmsconsignordernotify struct {

    // 订单商品信息
    
    OrderItem  *struct {
        Orderitemwlbwmsconsignordernotify  *Orderitemwlbwmsconsignordernotify `json:"orderitemwlbwmsconsignordernotify,omitempty"`
    } `json:"order_item,omitempty"`
    

}
*/

// Orderitemlistwlbwmsconsignordernotify 
type Orderitemlistwlbwmsconsignordernotify struct {

    // 订单商品信息
    OrderItem   *Orderitemwlbwmsconsignordernotify `json:"order_item,omitempty"`

}
