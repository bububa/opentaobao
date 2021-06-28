package tanx

// NativeTemplateAreaDto 
/* model for simplify = false
type NativeTemplateAreaDto struct {

    // 广告区域ID
    
    Id   int64 `json:"id,omitempty"`
    

    // 创意个数
    
    CreativeCount   int64 `json:"creative_count,omitempty"`
    

    // 创意要求
    
    Creative  *struct {
        NativeTemplateCreativeDto  *NativeTemplateCreativeDto `json:"native_template_creative_dto,omitempty"`
    } `json:"creative,omitempty"`
    

}
*/

// NativeTemplateAreaDto 
type NativeTemplateAreaDto struct {

    // 广告区域ID
    Id   int64 `json:"id,omitempty"`

    // 创意个数
    CreativeCount   int64 `json:"creative_count,omitempty"`

    // 创意要求
    Creative   *NativeTemplateCreativeDto `json:"creative,omitempty"`

}
