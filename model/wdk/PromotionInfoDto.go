package wdk

// PromotionInfoDTO 
type PromotionInfoDTO struct {
    // 展示文案
    DisplayText   string `json:"display_text,omitempty" xml:"display_text,omitempty"`
    // 优惠类型, 0:减钱,1:打折,2:一口价,3:组合优惠
    PromotionType   string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
    // 优惠金额，单位为分
    DiscountFee   int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
    // 创建时间
    CreateDateTime   string `json:"create_date_time,omitempty" xml:"create_date_time,omitempty"`
    // 活动ID
    ActivityId   int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
    // 活动名
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 活动开始时间
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
    // 限购信息
    LimitInfo   string `json:"limit_info,omitempty" xml:"limit_info,omitempty"`
    // 活动结束时间
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    // 活动类型, 1:单品活动,3:商品池活动
    ActivityType   int64 `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
}
