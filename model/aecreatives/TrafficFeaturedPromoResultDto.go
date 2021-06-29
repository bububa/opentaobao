package aecreatives

// TrafficFeaturedPromoResultDto 
type TrafficFeaturedPromoResultDto struct {

    // 当前返回数量
    
    CurrentRecordCount   int64 `json:"current_record_count,omitempty" xml:"current_record_count,omitempty"`
    

    // 返回主题活动列表
    
    Promos   []Promo `json:"promos,omitempty" xml:"promos,omitempty"`
    

}
