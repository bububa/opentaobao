package ascpchannel

// Outofstockitems 
type Outofstockitems struct {
    // 履约子单号
    SubOrderCode   string `json:"sub_order_code,omitempty" xml:"sub_order_code,omitempty"`
    // 货品ID
    ScItemId   string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
    // 货品缺货数量
    LackQuantity   int64 `json:"lack_quantity,omitempty" xml:"lack_quantity,omitempty"`
}
