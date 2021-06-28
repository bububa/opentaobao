package alsc

// ItemSelectedOpenInfoList 
/* model for simplify = false
type ItemSelectedOpenInfoList struct {

    // 规格id
    
    SkuId   string `json:"sku_id,omitempty"`
    

    // 菜单id
    
    DishId   string `json:"dish_id,omitempty"`
    

    // 规则号
    
    SkuOutNo   string `json:"sku_out_no,omitempty"`
    

    // 菜品号
    
    DishOutNo   string `json:"dish_out_no,omitempty"`
    

}
*/

// ItemSelectedOpenInfoList 
type ItemSelectedOpenInfoList struct {

    // 规格id
    SkuId   string `json:"sku_id,omitempty"`

    // 菜单id
    DishId   string `json:"dish_id,omitempty"`

    // 规则号
    SkuOutNo   string `json:"sku_out_no,omitempty"`

    // 菜品号
    DishOutNo   string `json:"dish_out_no,omitempty"`

}
