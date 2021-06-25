package wdk

// ItemDiscountActivityOperateRequest 
type ItemDiscountActivityOperateRequest struct {

    // 操作人ID（数字类型）
    CreatorId   string `json:"creator_id,omitempty"`

    // 操作人name
    CreatorName   string `json:"creator_name,omitempty"`

    // 活动ID
    ActId   int64 `json:"act_id,omitempty"`

    // 活动名称
    ActivityName   string `json:"activity_name,omitempty"`

    // 活动描述
    Description   string `json:"description,omitempty"`

    // 活动终端：1-APP，2-POS
    Terminals   []Number `json:"terminals,omitempty"`

    // 活动生效的经营店ID
    StoreIds   []String `json:"store_ids,omitempty"`

    // 外部活动ID（商家自定义）
    OutActId   string `json:"out_act_id,omitempty"`

    // 特价类型，1-减钱，2-一口价，3-打折
    DiscountType   int64 `json:"discount_type,omitempty"`

    // 活动开始时间
    StartTime   int64 `json:"start_time,omitempty"`

    // 活动结束时间
    EndTime   int64 `json:"end_time,omitempty"`

    // 活动人群编码，NEW_USER：新用户，OLD_USER：老用户，LIGHT_NEW_USER：闪购新客
    MemberCrowdCode   []String `json:"member_crowd_code,omitempty"`

}
