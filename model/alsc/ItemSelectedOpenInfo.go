package alsc

// ItemSelectedOpenInfo 
type ItemSelectedOpenInfo struct {

    // 规则ID
    SkuId   string `json:"sku_id,omitempty"`

    // 菜品ID
    DishId   string `json:"dish_id,omitempty"`

    // 外部菜品ID
    DishOutNo   string `json:"dish_out_no,omitempty"`

    // 外部规则ID
    SkuOutNo   string `json:"sku_out_no,omitempty"`

}
