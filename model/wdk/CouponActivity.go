package wdk

// CouponActivity 
/* model for simplify = false
type CouponActivity struct {

    // 优惠券适用范围 [rangeShop:店铺券;rangeItem:商品;rangeCategory:品类券]
    
    RangeType   string `json:"range_type,omitempty"`
    

    // 通用限购信息，-1为不限制，默认为不限制[如果同时设置了(每人活动期间总限领)和(每人每日限领)，则只能生效(每人活动期间总限领)]
    
    LimitInfo  *struct {
        LimitInfo  *LimitInfo `json:"limit_info,omitempty"`
    } `json:"limit_info,omitempty"`
    

    // 优惠适用场景[APP|POS|POS+APP分别对应的值为1|2|1,2]
    
    Terminals  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"terminals,omitempty"`
    

    // 商家活动id
    
    OutActId   string `json:"out_act_id,omitempty"`
    

    // 优惠券发放方式 [anonymous:匿名券;registered:记名券]
    
    SendType   string `json:"send_type,omitempty"`
    

    // 优惠券logo url，设置匿名券时为必传参数
    
    LogoUrl   string `json:"logo_url,omitempty"`
    

    // 参加活动的渠道店ids
    
    ShopIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"shop_ids,omitempty"`
    

    // 活动结束时间，时间戳[ms单位]
    
    EndTime   int64 `json:"end_time,omitempty"`
    

    // 适用类目
    
    ApplicableCategories  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"applicable_categories,omitempty"`
    

    // 券面额 [单位为分]
    
    Amount   int64 `json:"amount,omitempty"`
    

    // 活动开始时间，时间戳[ms单位]
    
    StartTime   int64 `json:"start_time,omitempty"`
    

    // 优惠券领取渠道 [fullSendCoupon:满额返券;fallingSendCoupon:天降红包;customerService:客服补偿发券;其他渠道请找接口人申请]
    
    ApplyChannels  struct {
        String  []string `json:"string,omitempty"`
    } `json:"apply_channels,omitempty"`
    

    // 领取后N日有效；如果设置了该值，则不需要设置优惠券的开始时间和结束时间 [有效期截止至领取日期+N天的23:59:59。例券设置有效期领取后5天有效，2018.1.1领取的券，有效期截止至2018.1.6 23:59:59； 例券设置有效期领取后0天有效，2018.1.1领取的券，有效期截止至2018.1.1 23:59:59]
    
    ValidDays   int64 `json:"valid_days,omitempty"`
    

    // 优惠券活动描述，不超过100个中文字符
    
    Description   string `json:"description,omitempty"`
    

    // [满减券特有，券类型为满减券时为必传参数]金额门槛，值为-1代表无门槛 [单位为分]
    
    StartFee   int64 `json:"start_fee,omitempty"`
    

    // 优惠券活动名称,不超过10个中文字符
    
    ActivityName   string `json:"activity_name,omitempty"`
    

    // 自定义的优惠券使用详情(支持多条)
    
    Details  struct {
        String  []string `json:"string,omitempty"`
    } `json:"details,omitempty"`
    

    // 优惠券优惠类型 [fullReduce:满减券;reduceTo:减至券，即一口价券]
    
    DiscountType   string `json:"discount_type,omitempty"`
    

    // [减至券特有]件数门槛，值为-1代表无门槛 [单位为整数]【已下线】
    
    StartCount   int64 `json:"start_count,omitempty"`
    

    // [减至券特有]优惠件数，限制最多优惠N件，值为-1代表不限制优惠件数 [单位为整数]【已下线】
    
    Count   int64 `json:"count,omitempty"`
    

}
*/

// CouponActivity 
type CouponActivity struct {

    // 优惠券适用范围 [rangeShop:店铺券;rangeItem:商品;rangeCategory:品类券]
    RangeType   string `json:"range_type,omitempty"`

    // 通用限购信息，-1为不限制，默认为不限制[如果同时设置了(每人活动期间总限领)和(每人每日限领)，则只能生效(每人活动期间总限领)]
    LimitInfo   *LimitInfo `json:"limit_info,omitempty"`

    // 优惠适用场景[APP|POS|POS+APP分别对应的值为1|2|1,2]
    Terminals   []int64 `json:"terminals,omitempty"`

    // 商家活动id
    OutActId   string `json:"out_act_id,omitempty"`

    // 优惠券发放方式 [anonymous:匿名券;registered:记名券]
    SendType   string `json:"send_type,omitempty"`

    // 优惠券logo url，设置匿名券时为必传参数
    LogoUrl   string `json:"logo_url,omitempty"`

    // 参加活动的渠道店ids
    ShopIds   []string `json:"shop_ids,omitempty"`

    // 活动结束时间，时间戳[ms单位]
    EndTime   int64 `json:"end_time,omitempty"`

    // 适用类目
    ApplicableCategories   []int64 `json:"applicable_categories,omitempty"`

    // 券面额 [单位为分]
    Amount   int64 `json:"amount,omitempty"`

    // 活动开始时间，时间戳[ms单位]
    StartTime   int64 `json:"start_time,omitempty"`

    // 优惠券领取渠道 [fullSendCoupon:满额返券;fallingSendCoupon:天降红包;customerService:客服补偿发券;其他渠道请找接口人申请]
    ApplyChannels   []string `json:"apply_channels,omitempty"`

    // 领取后N日有效；如果设置了该值，则不需要设置优惠券的开始时间和结束时间 [有效期截止至领取日期+N天的23:59:59。例券设置有效期领取后5天有效，2018.1.1领取的券，有效期截止至2018.1.6 23:59:59； 例券设置有效期领取后0天有效，2018.1.1领取的券，有效期截止至2018.1.1 23:59:59]
    ValidDays   int64 `json:"valid_days,omitempty"`

    // 优惠券活动描述，不超过100个中文字符
    Description   string `json:"description,omitempty"`

    // [满减券特有，券类型为满减券时为必传参数]金额门槛，值为-1代表无门槛 [单位为分]
    StartFee   int64 `json:"start_fee,omitempty"`

    // 优惠券活动名称,不超过10个中文字符
    ActivityName   string `json:"activity_name,omitempty"`

    // 自定义的优惠券使用详情(支持多条)
    Details   []string `json:"details,omitempty"`

    // 优惠券优惠类型 [fullReduce:满减券;reduceTo:减至券，即一口价券]
    DiscountType   string `json:"discount_type,omitempty"`

    // [减至券特有]件数门槛，值为-1代表无门槛 [单位为整数]【已下线】
    StartCount   int64 `json:"start_count,omitempty"`

    // [减至券特有]优惠件数，限制最多优惠N件，值为-1代表不限制优惠件数 [单位为整数]【已下线】
    Count   int64 `json:"count,omitempty"`

}
