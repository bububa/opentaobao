package opentrade

// GroupActivityResponse 
type GroupActivityResponse struct {
    // 组团活动id
    GroupActivityId   int64 `json:"group_activity_id,omitempty" xml:"group_activity_id,omitempty"`
    // 组团活动开始时间
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
    // 组团活动结束时间
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    // 成团有效期，单位为秒
    Expiration   int64 `json:"expiration,omitempty" xml:"expiration,omitempty"`
    // 成团的目标人数
    Goal   int64 `json:"goal,omitempty" xml:"goal,omitempty"`
    // 组团类型，0：拼团；1：团购
    GroupType   int64 `json:"group_type,omitempty" xml:"group_type,omitempty"`
    // 是否任何账号可开团。whitelist：仅白名单账号可开团 all：任何账号可开团
    AllowType   string `json:"allow_type,omitempty" xml:"allow_type,omitempty"`
    // 商品ID
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 组团购买的折扣价，单位为分，不能低于店铺最低折扣
    DiscountPrice   int64 `json:"discount_price,omitempty" xml:"discount_price,omitempty"`
    // 未成团处理办法，close：系统关单；continue：订单继续下行
    FailProcess   string `json:"fail_process,omitempty" xml:"fail_process,omitempty"`
    // 组团类型为团购，可限制团长针对一个商品的开团数量上限
    OpenLimit   int64 `json:"open_limit,omitempty" xml:"open_limit,omitempty"`
    // 允许开团的淘宝账号列表
    AllowWhiteList   []string `json:"allow_white_list,omitempty" xml:"allow_white_list>string,omitempty"`
}
