package qimen

// Detail 
/* model for simplify = false
type Detail struct {

    // 订单商品列表
    
    Items  struct {
        Item  []Item `json:"item,omitempty"`
    } `json:"items,omitempty"`
    

}
*/

// Detail 
type Detail struct {

    // 订单商品列表
    Items   []Item `json:"items,omitempty"`

}
