package tmallnr

// NrInventoryCheckDetailDto 
type NrInventoryCheckDetailDto struct {
    // 数量
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 商家的商品编码
    ScItemCode   string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
    // 淘宝后端商品的id号
    ScItemId   int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
    // 幂等值
    SubOrderId   string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
}
