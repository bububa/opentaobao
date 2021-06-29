package crm

// ExchangeActivityCreateDTO 
type ExchangeActivityCreateDTO struct {
    // 不包邮地区
    ExcludeArea   string `json:"exclude_area,omitempty" xml:"exclude_area,omitempty"`
    // 活动结束时间
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    // 商品是否包邮
    FreePostage   bool `json:"free_postage,omitempty" xml:"free_postage,omitempty"`
    // 优惠标签
    ActivityTag   string `json:"activity_tag,omitempty" xml:"activity_tag,omitempty"`
    // 活动开始时间
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
    // 商品ID
    ItmeId   int64 `json:"itme_id,omitempty" xml:"itme_id,omitempty"`
    // 活动名称
    ActivityName   string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
    // 商品一口价
    FixPrice   int64 `json:"fix_price,omitempty" xml:"fix_price,omitempty"`
}
