package promotion

// Meal 
/* model for simplify = false
type Meal struct {

    // 套餐id。
    
    MealId   int64 `json:"meal_id,omitempty"`
    

    // 搭配套餐名称。
    
    MealName   string `json:"meal_name,omitempty"`
    

    // 套餐一口价(单位是：分)
    
    MealPrice   float64 `json:"meal_price,omitempty"`
    

    // 搭配套餐商品列表。item_id为商品的id;item_show_name为商品显示名。因最多允许5个商品进行搭配，所以查询最多有5个，以json格式传出。
    
    ItemList   string `json:"item_list,omitempty"`
    

    // 普通运费模板id。若这个字段为空或0时，运费是卖家负责;若这个字段不为空，说明运费模板存在，运费是买家负责。
    
    PostageId   int64 `json:"postage_id,omitempty"`
    

    // 搭配套餐描述！
    
    MealMemo   string `json:"meal_memo,omitempty"`
    

    // 套餐状态。有效：VALID;失效：INVALID(有效套餐为可使用的套餐,无效套餐为套餐中有商品下架或库存为0时)。
    
    Status   string `json:"status,omitempty"`
    

}
*/

// Meal 
type Meal struct {

    // 套餐id。
    MealId   int64 `json:"meal_id,omitempty"`

    // 搭配套餐名称。
    MealName   string `json:"meal_name,omitempty"`

    // 套餐一口价(单位是：分)
    MealPrice   float64 `json:"meal_price,omitempty"`

    // 搭配套餐商品列表。item_id为商品的id;item_show_name为商品显示名。因最多允许5个商品进行搭配，所以查询最多有5个，以json格式传出。
    ItemList   string `json:"item_list,omitempty"`

    // 普通运费模板id。若这个字段为空或0时，运费是卖家负责;若这个字段不为空，说明运费模板存在，运费是买家负责。
    PostageId   int64 `json:"postage_id,omitempty"`

    // 搭配套餐描述！
    MealMemo   string `json:"meal_memo,omitempty"`

    // 套餐状态。有效：VALID;失效：INVALID(有效套餐为可使用的套餐,无效套餐为套餐中有商品下架或库存为0时)。
    Status   string `json:"status,omitempty"`

}
