package promotion

// CommonItemActivity 
type CommonItemActivity struct {

    // 优惠活动ID
    ActivityId   int64 `json:"activity_id,omitempty"`

    // 活动名称，不能超过32字符
    Name   string `json:"name,omitempty"`

    // 活动描述，不能超过100字符
    Description   string `json:"description,omitempty"`

    // 活动开始时间
    StartTime   string `json:"start_time,omitempty"`

    // 活动结束时间
    EndTime   string `json:"end_time,omitempty"`

    // 是否指定人群标签
    IsUserTag   bool `json:"is_user_tag,omitempty"`

    // 人群标签值
    UserTag   string `json:"user_tag,omitempty"`

}
