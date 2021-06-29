package eleenterpriseordernew

// FoodsInfo 
type FoodsInfo struct {

    // 餐品名称
    
    FoodName   string `json:"food_name,omitempty" xml:"food_name,omitempty"`
    

    // 餐品价格
    
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    

    // 餐品id
    
    FoodId   int64 `json:"food_id,omitempty" xml:"food_id,omitempty"`
    

    // 餐品数量
    
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
    

    // 规格Id
    
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    

}
