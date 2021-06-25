package promotion

// Detaillist 
type Detaillist struct {

    // 需要消耗星星数
    UnitPoint   string `json:"unit_point,omitempty"`

    // 价格分
    UnitPrice   int64 `json:"unit_price,omitempty"`

    // 数量
    Num   int64 `json:"num,omitempty"`

    // 扩展字段
    ExtMap   string `json:"ext_map,omitempty"`

}
