package alihealth2

// TopOrderGoodsDTO 
type TopOrderGoodsDTO struct {
    // 商品编码
    GoodsCode   string `json:"goods_code,omitempty" xml:"goods_code,omitempty"`
    // 商品数量
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
    // 商品名称
    GoodsName   string `json:"goods_name,omitempty" xml:"goods_name,omitempty"`
}
