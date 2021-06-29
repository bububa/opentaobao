package lstlogistics2

// SubOrders 
type SubOrders struct {
    // 外部商品编码
    OutItemCode   string `json:"out_item_code,omitempty" xml:"out_item_code,omitempty"`
    // 商品条码
    ItemBarCode   string `json:"item_bar_code,omitempty" xml:"item_bar_code,omitempty"`
    // 商品名称
    ItemName   string `json:"item_name,omitempty" xml:"item_name,omitempty"`
    // 零售通子订单号
    LstSubOrderId   string `json:"lst_sub_order_id,omitempty" xml:"lst_sub_order_id,omitempty"`
    // 商品数量
    ItemQuantity   int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
    // 揽收数量
    PickQuantity   int64 `json:"pick_quantity,omitempty" xml:"pick_quantity,omitempty"`
    // 外部子订单号
    OutSubOrderId   string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
    // 签收数量
    SignQuantity   int64 `json:"sign_quantity,omitempty" xml:"sign_quantity,omitempty"`
}
