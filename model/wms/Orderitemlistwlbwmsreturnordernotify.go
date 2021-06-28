package wms

// Orderitemlistwlbwmsreturnordernotify 
/* model for simplify = false
type Orderitemlistwlbwmsreturnordernotify struct {

    // 1
    
    OrderItem  *struct {
        Orderitemwlbwmsreturnordernotify  *Orderitemwlbwmsreturnordernotify `json:"orderitemwlbwmsreturnordernotify,omitempty"`
    } `json:"order_item,omitempty"`
    

}
*/

// Orderitemlistwlbwmsreturnordernotify 
type Orderitemlistwlbwmsreturnordernotify struct {

    // 1
    OrderItem   *Orderitemwlbwmsreturnordernotify `json:"order_item,omitempty"`

}
