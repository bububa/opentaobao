package hotel

// HotelTopDetailsParam 
/* model for simplify = false
type HotelTopDetailsParam struct {

    // 成人数
    
    AdultNum   int64 `json:"adult_num,omitempty"`
    

    // 入住日期
    
    CheckIn   string `json:"check_in,omitempty"`
    

    // 离店日期
    
    CheckOut   string `json:"check_out,omitempty"`
    

    // 儿童年龄
    
    ChildrenAges   string `json:"children_ages,omitempty"`
    

    // 城市code
    
    CityCode   string `json:"city_code,omitempty"`
    

    // 筛选
    
    ScreenCodes   string `json:"screen_codes,omitempty"`
    

    // 卖家id
    
    SellerIds  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"seller_ids,omitempty"`
    

    // 标准酒店id
    
    Shid   int64 `json:"shid,omitempty"`
    

    // 渠道ttid
    
    Ttid   string `json:"ttid,omitempty"`
    

    // 用户id
    
    UserId   int64 `json:"user_id,omitempty"`
    

}
*/

// HotelTopDetailsParam 
type HotelTopDetailsParam struct {

    // 成人数
    AdultNum   int64 `json:"adult_num,omitempty"`

    // 入住日期
    CheckIn   string `json:"check_in,omitempty"`

    // 离店日期
    CheckOut   string `json:"check_out,omitempty"`

    // 儿童年龄
    ChildrenAges   string `json:"children_ages,omitempty"`

    // 城市code
    CityCode   string `json:"city_code,omitempty"`

    // 筛选
    ScreenCodes   string `json:"screen_codes,omitempty"`

    // 卖家id
    SellerIds   []int64 `json:"seller_ids,omitempty"`

    // 标准酒店id
    Shid   int64 `json:"shid,omitempty"`

    // 渠道ttid
    Ttid   string `json:"ttid,omitempty"`

    // 用户id
    UserId   int64 `json:"user_id,omitempty"`

}
