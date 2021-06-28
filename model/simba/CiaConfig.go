package simba

// CiaConfig 
/* model for simplify = false
type CiaConfig struct {

    // 目标点击量
    
    TargetClick   int64 `json:"target_click,omitempty"`
    

    // 出价偏好（3:促进收藏加购，4:促进点击，5:促进成交）
    
    BidTargetType   int64 `json:"bid_target_type,omitempty"`
    

    // 最高溢价比例
    
    MaxPremium   int64 `json:"max_premium,omitempty"`
    

    // 是否自动流转
    
    IsCirculation   bool `json:"is_circulation,omitempty"`
    

    // 是否开启智能出价
    
    IsSmartBidding   bool `json:"is_smart_bidding,omitempty"`
    

    // 推广组id
    
    AdGroupId   int64 `json:"ad_group_id,omitempty"`
    

}
*/

// CiaConfig 
type CiaConfig struct {

    // 目标点击量
    TargetClick   int64 `json:"target_click,omitempty"`

    // 出价偏好（3:促进收藏加购，4:促进点击，5:促进成交）
    BidTargetType   int64 `json:"bid_target_type,omitempty"`

    // 最高溢价比例
    MaxPremium   int64 `json:"max_premium,omitempty"`

    // 是否自动流转
    IsCirculation   bool `json:"is_circulation,omitempty"`

    // 是否开启智能出价
    IsSmartBidding   bool `json:"is_smart_bidding,omitempty"`

    // 推广组id
    AdGroupId   int64 `json:"ad_group_id,omitempty"`

}
