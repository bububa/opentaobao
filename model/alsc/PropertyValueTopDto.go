package alsc

// PropertyValueTopDTO 
type PropertyValueTopDTO struct {
    // 属性值
    Value   string `json:"value,omitempty" xml:"value,omitempty"`
    // 属性ID
    PropertyId   int64 `json:"property_id,omitempty" xml:"property_id,omitempty"`
}
