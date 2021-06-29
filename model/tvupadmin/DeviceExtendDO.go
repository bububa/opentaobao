package tvupadmin

// DeviceExtendDO 
type DeviceExtendDO struct {
    // 牌照方
    Bcp   int64 `json:"bcp,omitempty" xml:"bcp,omitempty"`
    // 设备型号
    DeviceModel   string `json:"device_model,omitempty" xml:"device_model,omitempty"`
    // 设备内部型号
    InnerDeviceModel   string `json:"inner_device_model,omitempty" xml:"inner_device_model,omitempty"`
    // 设备类型
    TerminalType   string `json:"terminal_type,omitempty" xml:"terminal_type,omitempty"`
    // 厂商名称
    BrandName   string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
    // 厂商ID
    BrandId   int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    // 创建时间
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`
}
