package miniappopen

// DistributionOrderSaveOpenRequest 
type DistributionOrderSaveOpenRequest struct {

    // 投放的小程序链接
    
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
    

    // 时效类型 0:自定义时效； 1：订购期有效；为0时，开始时间和结束时间必传
    
    TimeType   string `json:"time_type,omitempty" xml:"time_type,omitempty"`
    

    // 投放的小程序id
    
    AppId   int64 `json:"app_id,omitempty" xml:"app_id,omitempty"`
    

    // 要投放的卡片id
    
    CardId   int64 `json:"card_id,omitempty" xml:"card_id,omitempty"`
    

    // 投放计划名字
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 要投放的场景id
    
    SceneId   int64 `json:"scene_id,omitempty" xml:"scene_id,omitempty"`
    

    // 投放计划开始时间
    
    StartTime   int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
    

    // 投放计划结束时间
    
    EndTime   int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
    

}
