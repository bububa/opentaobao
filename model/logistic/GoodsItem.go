package logistic

// GoodsItem 
type GoodsItem struct {

    // 子订单编号
    Oid   int64 `json:"oid,omitempty"`

    // 货品数量
    Num   int64 `json:"num,omitempty"`

    // 货品ID
    GoodsId   string `json:"goods_id,omitempty"`

}
