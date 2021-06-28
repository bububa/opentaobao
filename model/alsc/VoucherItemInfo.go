package alsc

// VoucherItemInfo 
type VoucherItemInfo struct {

    // sku规格id
    
    SkuId   string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    

    // 菜品id
    
    DishId   string `json:"dish_id,omitempty" xml:"dish_id,omitempty"`
    

    // 外部sku规格id
    
    SkuOutNo   string `json:"sku_out_no,omitempty" xml:"sku_out_no,omitempty"`
    

    // 外部菜品id
    
    DishOutNo   string `json:"dish_out_no,omitempty" xml:"dish_out_no,omitempty"`
    

}
