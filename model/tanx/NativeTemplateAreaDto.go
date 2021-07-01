package tanx

// NativeTemplateAreaDto 
type NativeTemplateAreaDto struct {
    // 广告区域ID
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 创意个数
    CreativeCount   int64 `json:"creative_count,omitempty" xml:"creative_count,omitempty"`
    // 创意要求
    Creative   *NativeTemplateCreativeDto `json:"creative,omitempty" xml:"creative,omitempty"`
}
