package rhino

// SkuBo 
type SkuBo struct {
    // 款式id
    StyleId   string `json:"style_id,omitempty" xml:"style_id,omitempty"`
    // 工艺版本id
    TechnologyVersion   string `json:"technology_version,omitempty" xml:"technology_version,omitempty"`
    // 颜色id
    ColorId   int64 `json:"color_id,omitempty" xml:"color_id,omitempty"`
    // 尺码id
    SizeId   int64 `json:"size_id,omitempty" xml:"size_id,omitempty"`
}
