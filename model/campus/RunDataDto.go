package campus

// RunDataDto 
type RunDataDto struct {

    // 参数code
    
    PropertyCode   string `json:"property_code,omitempty" xml:"property_code,omitempty"`
    

    // 参数名称
    
    PropertyName   string `json:"property_name,omitempty" xml:"property_name,omitempty"`
    

    // 实时参数值
    
    Value   string `json:"value,omitempty" xml:"value,omitempty"`
    

    // 参数在数据字典所属分类，如30600，代表character
    
    PropertyKind   int64 `json:"property_kind,omitempty" xml:"property_kind,omitempty"`
    

    // 单位编码
    
    UnitCode   string `json:"unit_code,omitempty" xml:"unit_code,omitempty"`
    

    // 被关联物理设备编码
    
    RefDeviceId   string `json:"ref_device_id,omitempty" xml:"ref_device_id,omitempty"`
    

    // 被关联参数点属性
    
    RefPropertyCode   string `json:"ref_property_code,omitempty" xml:"ref_property_code,omitempty"`
    

}
