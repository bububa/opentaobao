package aetask

// QueryDeliveryRequestDto 
type QueryDeliveryRequestDto struct {

    // 0:不展示预热 1：展示预热
    
    PreDisplay   int64 `json:"pre_display,omitempty" xml:"pre_display,omitempty"`
    

    // 投放场景id
    
    SceneId   int64 `json:"scene_id,omitempty" xml:"scene_id,omitempty"`
    

    // 用户版本信息
    
    Ttid   string `json:"ttid,omitempty" xml:"ttid,omitempty"`
    

    // 语言
    
    Language   string `json:"language,omitempty" xml:"language,omitempty"`
    

    // 国家
    
    Country   string `json:"country,omitempty" xml:"country,omitempty"`
    

    // 准入key
    
    ProjectAppKey   string `json:"project_app_key,omitempty" xml:"project_app_key,omitempty"`
    

}
