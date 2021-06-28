package gameact

// ActivityVO 
/* model for simplify = false
type ActivityVO struct {

    // 活动连接
    
    ActivityUrl   string `json:"activity_url,omitempty"`
    

    // 活动id
    
    ActivityId   int64 `json:"activity_id,omitempty"`
    

    // 活动名称
    
    Name   string `json:"name,omitempty"`
    

    // 1970年到现在的毫秒数
    
    StartTime   int64 `json:"start_time,omitempty"`
    

    // 1970年距离现在的毫秒数
    
    EndTime   int64 `json:"end_time,omitempty"`
    

    // 活动描述
    
    Description   string `json:"description,omitempty"`
    

    // 奖项列表
    
    Awards  struct {
        AwardVO  []AwardVO `json:"award_vo,omitempty"`
    } `json:"awards,omitempty"`
    

    // 运营和cp约定的唯一事件标示
    
    EventKey   string `json:"event_key,omitempty"`
    

    // 积分/金牌消耗
    
    ConsumeAmount   int64 `json:"consume_amount,omitempty"`
    

    // 抽奖类型
    
    LuckyType   int64 `json:"lucky_type,omitempty"`
    

    // 抽奖渠道
    
    LuckyChannel   int64 `json:"lucky_channel,omitempty"`
    

    // 抽奖次数（免费）
    
    AccessAmount   int64 `json:"access_amount,omitempty"`
    

}
*/

// ActivityVO 
type ActivityVO struct {

    // 活动连接
    ActivityUrl   string `json:"activity_url,omitempty"`

    // 活动id
    ActivityId   int64 `json:"activity_id,omitempty"`

    // 活动名称
    Name   string `json:"name,omitempty"`

    // 1970年到现在的毫秒数
    StartTime   int64 `json:"start_time,omitempty"`

    // 1970年距离现在的毫秒数
    EndTime   int64 `json:"end_time,omitempty"`

    // 活动描述
    Description   string `json:"description,omitempty"`

    // 奖项列表
    Awards   []AwardVO `json:"awards,omitempty"`

    // 运营和cp约定的唯一事件标示
    EventKey   string `json:"event_key,omitempty"`

    // 积分/金牌消耗
    ConsumeAmount   int64 `json:"consume_amount,omitempty"`

    // 抽奖类型
    LuckyType   int64 `json:"lucky_type,omitempty"`

    // 抽奖渠道
    LuckyChannel   int64 `json:"lucky_channel,omitempty"`

    // 抽奖次数（免费）
    AccessAmount   int64 `json:"access_amount,omitempty"`

}
