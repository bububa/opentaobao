package campus

// DeviceDataApiDto 
type DeviceDataApiDto struct {

    // 设备id
    
    DeviceId   string `json:"device_id,omitempty" xml:"device_id,omitempty"`
    

    // 设备code
    
    DeviceCode   string `json:"device_code,omitempty" xml:"device_code,omitempty"`
    

    // 参数点code
    
    PropertyCode   string `json:"property_code,omitempty" xml:"property_code,omitempty"`
    

    // 历史数据值
    
    Value   string `json:"value,omitempty" xml:"value,omitempty"`
    

    // 时间戳
    
    Timestamp   int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
    

    // 值类型id
    
    ValueType   int64 `json:"value_type,omitempty" xml:"value_type,omitempty"`
    

    // 值类型名称
    
    ValueTypeName   string `json:"value_type_name,omitempty" xml:"value_type_name,omitempty"`
    

    // 单位id
    
    UnitId   int64 `json:"unit_id,omitempty" xml:"unit_id,omitempty"`
    

    // 单位编码
    
    UnitCode   string `json:"unit_code,omitempty" xml:"unit_code,omitempty"`
    

}
