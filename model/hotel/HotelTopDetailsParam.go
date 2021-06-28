package hotel

// HotelTopDetailsParam 
type HotelTopDetailsParam struct {

    // 成人数
    
    AdultNum   int64 `json:"adult_num,omitempty" xml:"adult_num,omitempty"`
    

    // 入住日期
    
    CheckIn   string `json:"check_in,omitempty" xml:"check_in,omitempty"`
    

    // 离店日期
    
    CheckOut   string `json:"check_out,omitempty" xml:"check_out,omitempty"`
    

    // 儿童年龄
    
    ChildrenAges   string `json:"children_ages,omitempty" xml:"children_ages,omitempty"`
    

    // 城市code
    
    CityCode   string `json:"city_code,omitempty" xml:"city_code,omitempty"`
    

    // 筛选
    
    ScreenCodes   string `json:"screen_codes,omitempty" xml:"screen_codes,omitempty"`
    

    // 卖家id
    
    SellerIds   []int64 `json:"seller_ids,omitempty" xml:"seller_ids>int64,omitempty"`
    

    // 标准酒店id
    
    Shid   int64 `json:"shid,omitempty" xml:"shid,omitempty"`
    

    // 渠道ttid
    
    Ttid   string `json:"ttid,omitempty" xml:"ttid,omitempty"`
    

    // 用户id
    
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
    

}
