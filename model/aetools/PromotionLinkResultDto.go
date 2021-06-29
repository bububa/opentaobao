package aetools

// PromotionLinkResultDto 
type PromotionLinkResultDto struct {

    // 推广链接列表
    
    PromotionLinks   []PromotionLink `json:"promotion_links,omitempty" xml:"promotion_links,omitempty"`
    

    // 返回总量
    
    TotalResultCount   int64 `json:"total_result_count,omitempty" xml:"total_result_count,omitempty"`
    

    // 推广者TrackingId
    
    TrackingId   string `json:"tracking_id,omitempty" xml:"tracking_id,omitempty"`
    

}
