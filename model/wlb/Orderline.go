package wlb

// Orderline 
type Orderline struct {

    // 订单行号
    OrderLineNo   string `json:"order_line_no,omitempty"`

    // 商品ID
    ItemId   string `json:"item_id,omitempty"`

    // 商品编码
    ItemCode   string `json:"item_code,omitempty"`

    // 数量
    Amount   string `json:"amount,omitempty"`

}
