package interact

// InteractDTO 
/* model for simplify = false
type InteractDTO struct {

    // 互动开始时间
    
    StartTime   string `json:"start_time,omitempty"`
    

    // 互动游戏app name
    
    AppName   string `json:"app_name,omitempty"`
    

    // 店铺ID
    
    ShopId   int64 `json:"shop_id,omitempty"`
    

    // 卖家ID
    
    SellerId   int64 `json:"seller_id,omitempty"`
    

    // 互动描述
    
    Description   string `json:"description,omitempty"`
    

    // 互动实例ID
    
    InteractId   string `json:"interact_id,omitempty"`
    

    // 互动结束时间
    
    EndTime   string `json:"end_time,omitempty"`
    

    // 互动游戏app key
    
    AppKey   string `json:"app_key,omitempty"`
    

    // 互动实例名称
    
    InstanceName   string `json:"instance_name,omitempty"`
    

}
*/

// InteractDTO 
type InteractDTO struct {

    // 互动开始时间
    StartTime   string `json:"start_time,omitempty"`

    // 互动游戏app name
    AppName   string `json:"app_name,omitempty"`

    // 店铺ID
    ShopId   int64 `json:"shop_id,omitempty"`

    // 卖家ID
    SellerId   int64 `json:"seller_id,omitempty"`

    // 互动描述
    Description   string `json:"description,omitempty"`

    // 互动实例ID
    InteractId   string `json:"interact_id,omitempty"`

    // 互动结束时间
    EndTime   string `json:"end_time,omitempty"`

    // 互动游戏app key
    AppKey   string `json:"app_key,omitempty"`

    // 互动实例名称
    InstanceName   string `json:"instance_name,omitempty"`

}
