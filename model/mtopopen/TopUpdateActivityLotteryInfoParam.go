package mtopopen

// TopUpdateActivityLotteryInfoParam 
/* model for simplify = false
type TopUpdateActivityLotteryInfoParam struct {

    // 商家通过isv创建的活动id
    
    ActivityBizId   string `json:"activity_biz_id,omitempty"`
    

    // isv的appkey
    
    AppKey   string `json:"app_key,omitempty"`
    

    // 可随意传，目前没有用到，扩展字段
    
    BannerUrl   string `json:"banner_url,omitempty"`
    

    // 商家创建的红包总额
    
    BenefitAmount   string `json:"benefit_amount,omitempty"`
    

    // 可随意传，目前没有用到，扩展字段
    
    BenefitAttribute   string `json:"benefit_attribute,omitempty"`
    

    // 商家创建的红包面额
    
    BenefitDenomination   int64 `json:"benefit_denomination,omitempty"`
    

    // 商家选择使用的红包权益id
    
    BenefitId   int64 `json:"benefit_id,omitempty"`
    

    // 目前填1，代表支付宝红包
    
    BenefitType   string `json:"benefit_type,omitempty"`
    

    // 目前必须填4，代表互动城
    
    BizType   string `json:"biz_type,omitempty"`
    

    // 商家创建的isv活动的结束时间
    
    EndTime   string `json:"end_time,omitempty"`
    

    // 抽奖活动开始时间，long型
    
    LotteryActivityEndDate   int64 `json:"lottery_activity_end_date,omitempty"`
    

    // 抽奖活动结束时间
    
    LotteryActivityStartDate   int64 `json:"lottery_activity_start_date,omitempty"`
    

    // 抽奖活动名称
    
    Name   string `json:"name,omitempty"`
    

    // 奖品扩展字段
    
    PrizeExtAttribute   string `json:"prize_ext_attribute,omitempty"`
    

    // 1元红包
    
    PrizeName   string `json:"prize_name,omitempty"`
    

    // 奖品扩展字段
    
    PrizeParamAttribute   string `json:"prize_param_attribute,omitempty"`
    

    // 奖品总数
    
    PrizeQuantity   string `json:"prize_quantity,omitempty"`
    

    // 奖品总数
    
    PrizeRemainQuantity   string `json:"prize_remain_quantity,omitempty"`
    

    // 必须填1，代表红包
    
    PrizeType   string `json:"prize_type,omitempty"`
    

    // 0～1之间，代表奖池的概率
    
    Probability   string `json:"probability,omitempty"`
    

    // 用户活动期间总共中奖次数限制
    
    WinPermissionActivityCount   int64 `json:"win_permission_activity_count,omitempty"`
    

    // 用户每天总共中奖次数限制
    
    WinPermissionDayCount   int64 `json:"win_permission_day_count,omitempty"`
    

    // 商家创建的isv活动的开始时间
    
    StartTime   string `json:"start_time,omitempty"`
    

}
*/

// TopUpdateActivityLotteryInfoParam 
type TopUpdateActivityLotteryInfoParam struct {

    // 商家通过isv创建的活动id
    ActivityBizId   string `json:"activity_biz_id,omitempty"`

    // isv的appkey
    AppKey   string `json:"app_key,omitempty"`

    // 可随意传，目前没有用到，扩展字段
    BannerUrl   string `json:"banner_url,omitempty"`

    // 商家创建的红包总额
    BenefitAmount   string `json:"benefit_amount,omitempty"`

    // 可随意传，目前没有用到，扩展字段
    BenefitAttribute   string `json:"benefit_attribute,omitempty"`

    // 商家创建的红包面额
    BenefitDenomination   int64 `json:"benefit_denomination,omitempty"`

    // 商家选择使用的红包权益id
    BenefitId   int64 `json:"benefit_id,omitempty"`

    // 目前填1，代表支付宝红包
    BenefitType   string `json:"benefit_type,omitempty"`

    // 目前必须填4，代表互动城
    BizType   string `json:"biz_type,omitempty"`

    // 商家创建的isv活动的结束时间
    EndTime   string `json:"end_time,omitempty"`

    // 抽奖活动开始时间，long型
    LotteryActivityEndDate   int64 `json:"lottery_activity_end_date,omitempty"`

    // 抽奖活动结束时间
    LotteryActivityStartDate   int64 `json:"lottery_activity_start_date,omitempty"`

    // 抽奖活动名称
    Name   string `json:"name,omitempty"`

    // 奖品扩展字段
    PrizeExtAttribute   string `json:"prize_ext_attribute,omitempty"`

    // 1元红包
    PrizeName   string `json:"prize_name,omitempty"`

    // 奖品扩展字段
    PrizeParamAttribute   string `json:"prize_param_attribute,omitempty"`

    // 奖品总数
    PrizeQuantity   string `json:"prize_quantity,omitempty"`

    // 奖品总数
    PrizeRemainQuantity   string `json:"prize_remain_quantity,omitempty"`

    // 必须填1，代表红包
    PrizeType   string `json:"prize_type,omitempty"`

    // 0～1之间，代表奖池的概率
    Probability   string `json:"probability,omitempty"`

    // 用户活动期间总共中奖次数限制
    WinPermissionActivityCount   int64 `json:"win_permission_activity_count,omitempty"`

    // 用户每天总共中奖次数限制
    WinPermissionDayCount   int64 `json:"win_permission_day_count,omitempty"`

    // 商家创建的isv活动的开始时间
    StartTime   string `json:"start_time,omitempty"`

}
