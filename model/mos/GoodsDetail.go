package mos

// GoodsDetail 
type GoodsDetail struct {
    // 商户自有的专柜号
    ShopNo   string `json:"shop_no,omitempty" xml:"shop_no,omitempty"`
    // 专柜名
    ShopName   string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
    // 商户自有的商品ID
    GoodsId   string `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
    // 商品名称
    GoodsName   string `json:"goods_name,omitempty" xml:"goods_name,omitempty"`
    // 商品总价，单位为分。与商品单价之间是二选一的关系。可以为负值，用于表达折扣
    Amount   string `json:"amount,omitempty" xml:"amount,omitempty"`
    // 商品单价，单位为分。与商品总价之间是二选一的关系。可以为负值，用于表达折扣
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    // 商品数量，支持小数
    Quantity   string `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
