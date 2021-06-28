package interact

// OpenMealDo 
/* model for simplify = false
type OpenMealDo struct {

    // 手淘购买链接
    
    H5BuyUrl   string `json:"h5_buy_url,omitempty"`
    

    // 单位分，套餐总价
    
    MealPrice   int64 `json:"meal_price,omitempty"`
    

    // 套餐商品列表
    
    Items  struct {
        OpenMealItemDo  []OpenMealItemDo `json:"open_meal_item_do,omitempty"`
    } `json:"items,omitempty"`
    

    // 套餐名称
    
    MealName   string `json:"meal_name,omitempty"`
    

    // 套餐开始时间戳
    
    StartTime   string `json:"start_time,omitempty"`
    

    // 套餐结束时间戳
    
    EndTime   string `json:"end_time,omitempty"`
    

}
*/

// OpenMealDo 
type OpenMealDo struct {

    // 手淘购买链接
    H5BuyUrl   string `json:"h5_buy_url,omitempty"`

    // 单位分，套餐总价
    MealPrice   int64 `json:"meal_price,omitempty"`

    // 套餐商品列表
    Items   []OpenMealItemDo `json:"items,omitempty"`

    // 套餐名称
    MealName   string `json:"meal_name,omitempty"`

    // 套餐开始时间戳
    StartTime   string `json:"start_time,omitempty"`

    // 套餐结束时间戳
    EndTime   string `json:"end_time,omitempty"`

}
