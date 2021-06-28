package caipiao

// WangcaiMarketingDetail 
/* model for simplify = false
type WangcaiMarketingDetail struct {

    // 外部系统主键
    
    BizId   string `json:"biz_id,omitempty"`
    

    // 活动ID,一个活动可以关联多个送彩票设置
    
    ActivityId   string `json:"activity_id,omitempty"`
    

    // 活动名称
    
    ActivityName   string `json:"activity_name,omitempty"`
    

    // 活动类型,0全店/1指定商品
    
    ActivityType   int64 `json:"activity_type,omitempty"`
    

    // 活动开始时间
    
    BeginTime   string `json:"begin_time,omitempty"`
    

    // 活动结束时间
    
    EndTime   string `json:"end_time,omitempty"`
    

    // 活动的最小金额门槛,以分为单位
    
    MinAmount   int64 `json:"min_amount,omitempty"`
    

    // 赠送的彩种,1双色球/8大乐透
    
    LotteryTypeId   int64 `json:"lottery_type_id,omitempty"`
    

    // 赠送的彩票注数
    
    Quantity   int64 `json:"quantity,omitempty"`
    

    // 参与活动的商品ID
    
    Items  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"items,omitempty"`
    

}
*/

// WangcaiMarketingDetail 
type WangcaiMarketingDetail struct {

    // 外部系统主键
    BizId   string `json:"biz_id,omitempty"`

    // 活动ID,一个活动可以关联多个送彩票设置
    ActivityId   string `json:"activity_id,omitempty"`

    // 活动名称
    ActivityName   string `json:"activity_name,omitempty"`

    // 活动类型,0全店/1指定商品
    ActivityType   int64 `json:"activity_type,omitempty"`

    // 活动开始时间
    BeginTime   string `json:"begin_time,omitempty"`

    // 活动结束时间
    EndTime   string `json:"end_time,omitempty"`

    // 活动的最小金额门槛,以分为单位
    MinAmount   int64 `json:"min_amount,omitempty"`

    // 赠送的彩种,1双色球/8大乐透
    LotteryTypeId   int64 `json:"lottery_type_id,omitempty"`

    // 赠送的彩票注数
    Quantity   int64 `json:"quantity,omitempty"`

    // 参与活动的商品ID
    Items   []int64 `json:"items,omitempty"`

}
