package tbk

// LocalizationMapData 
/* model for simplify = false
type LocalizationMapData struct {

    // 本地化-配送时间
    
    OrderLeadTime   string `json:"order_lead_time,omitempty"`
    

    // 本地化-用户评分
    
    UserRating   string `json:"user_rating,omitempty"`
    

    // 本地化-起送价
    
    DeliveryMinPrice   string `json:"delivery_min_price,omitempty"`
    

    // 本地化-配送费
    
    DeliveryFee   string `json:"delivery_fee,omitempty"`
    

    // 本地化-配送费原价
    
    OriginalDeliveryFee   string `json:"original_delivery_fee,omitempty"`
    

    // 本地化-配送类型；0：蜂鸟专送 1：蜂鸟快送 2：商家自配 3: 全城送
    
    DeliveryType   string `json:"delivery_type,omitempty"`
    

    // 本地化-推荐理由
    
    RecommendReasons  struct {
        String  []string `json:"string,omitempty"`
    } `json:"recommend_reasons,omitempty"`
    

    // 本地化-营销标签
    
    SaleTags  struct {
        String  []string `json:"string,omitempty"`
    } `json:"sale_tags,omitempty"`
    

    // 本地化-单店商品列表
    
    FoodItemList  struct {
        FoodMapData  []FoodMapData `json:"food_map_data,omitempty"`
    } `json:"food_item_list,omitempty"`
    

}
*/

// LocalizationMapData 
type LocalizationMapData struct {

    // 本地化-配送时间
    OrderLeadTime   string `json:"order_lead_time,omitempty"`

    // 本地化-用户评分
    UserRating   string `json:"user_rating,omitempty"`

    // 本地化-起送价
    DeliveryMinPrice   string `json:"delivery_min_price,omitempty"`

    // 本地化-配送费
    DeliveryFee   string `json:"delivery_fee,omitempty"`

    // 本地化-配送费原价
    OriginalDeliveryFee   string `json:"original_delivery_fee,omitempty"`

    // 本地化-配送类型；0：蜂鸟专送 1：蜂鸟快送 2：商家自配 3: 全城送
    DeliveryType   string `json:"delivery_type,omitempty"`

    // 本地化-推荐理由
    RecommendReasons   []string `json:"recommend_reasons,omitempty"`

    // 本地化-营销标签
    SaleTags   []string `json:"sale_tags,omitempty"`

    // 本地化-单店商品列表
    FoodItemList   []FoodMapData `json:"food_item_list,omitempty"`

}
