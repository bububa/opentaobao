package wdkitem

// MzPromotionDto 
type MzPromotionDto struct {

    // 活动ID
    
    ActId   int64 `json:"act_id,omitempty" xml:"act_id,omitempty"`
    

    // 活动名称
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 买赠活动文案
    
    MzDisplayText   string `json:"mz_display_text,omitempty" xml:"mz_display_text,omitempty"`
    

    // 活动开始时间
    
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
    

    // 活动结束时间
    
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    

}
