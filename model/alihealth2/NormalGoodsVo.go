package alihealth2

// NormalGoodsVo 
type NormalGoodsVo struct {
    // 标品ID
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 标品名称
    ItemName   string `json:"item_name,omitempty" xml:"item_name,omitempty"`
    // 成本价，单位元
    CostPrice   string `json:"cost_price,omitempty" xml:"cost_price,omitempty"`
    // 售价，单位元
    SoldPrice   string `json:"sold_price,omitempty" xml:"sold_price,omitempty"`
}
