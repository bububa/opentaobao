package nrt

// CategoryPropDto 
type CategoryPropDto struct {

    // 属性ID
    
    PropertyId   int64 `json:"property_id,omitempty" xml:"property_id,omitempty"`
    

    // 属性名
    
    PropertyName   string `json:"property_name,omitempty" xml:"property_name,omitempty"`
    

    // 值ID
    
    ValueId   int64 `json:"value_id,omitempty" xml:"value_id,omitempty"`
    

    // 值名
    
    ValueText   string `json:"value_text,omitempty" xml:"value_text,omitempty"`
    

}
