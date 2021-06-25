package mos

// CodeGoodsDto 
type CodeGoodsDto struct {

    // 数量
    Quantity   string `json:"quantity,omitempty"`

    // id
    GoodsId   int64 `json:"goods_id,omitempty"`

    // zi单号
    SubNo   string `json:"sub_no,omitempty"`

}
