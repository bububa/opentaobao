package tmallgenie

// RetValue 
type RetValue struct {

    // 设备图片
    
    DeviceIcon   string `json:"device_icon,omitempty" xml:"device_icon,omitempty"`
    

    // 型号
    
    DeviceModel   string `json:"device_model,omitempty" xml:"device_model,omitempty"`
    

    // 别名
    
    DeviceAlias   string `json:"device_alias,omitempty" xml:"device_alias,omitempty"`
    

    // 原本设备类型
    
    OrginDeviceCategory   string `json:"orgin_device_category,omitempty" xml:"orgin_device_category,omitempty"`
    

    // 可控制的属性
    
    DeviceProperties   string `json:"device_properties,omitempty" xml:"device_properties,omitempty"`
    

    // 详细型号
    
    ParticularModel   string `json:"particular_model,omitempty" xml:"particular_model,omitempty"`
    

    // 设备类型英文
    
    DeviceCategoryEn   string `json:"device_category_en,omitempty" xml:"device_category_en,omitempty"`
    

    // 设备类型中文
    
    DeviceCategory   string `json:"device_category,omitempty" xml:"device_category,omitempty"`
    

    // 品牌
    
    DeviceBrand   string `json:"device_brand,omitempty" xml:"device_brand,omitempty"`
    

    // 设备id
    
    DeviceId   string `json:"device_id,omitempty" xml:"device_id,omitempty"`
    

    // 位置
    
    DeviceZone   string `json:"device_zone,omitempty" xml:"device_zone,omitempty"`
    

}
