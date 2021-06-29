package alsc

// MenuDetailOpenInfo 
type MenuDetailOpenInfo struct {
    // 折扣率
    ProDiscount   string `json:"pro_discount,omitempty" xml:"pro_discount,omitempty"`
    // 0-折扣，1-固定价
    ProType   string `json:"pro_type,omitempty" xml:"pro_type,omitempty"`
    // 规格ID
    SkuId   string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    // 外部规菜品D
    DishOutNo   string `json:"dish_out_no,omitempty" xml:"dish_out_no,omitempty"`
    // 外部规则ID
    SkuOutNo   string `json:"sku_out_no,omitempty" xml:"sku_out_no,omitempty"`
    // 菜品ID
    DishId   string `json:"dish_id,omitempty" xml:"dish_id,omitempty"`
    // 固定价
    ProPrice   int64 `json:"pro_price,omitempty" xml:"pro_price,omitempty"`
}
