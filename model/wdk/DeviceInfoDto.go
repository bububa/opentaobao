package wdk

// DeviceInfoDTO 
type DeviceInfoDTO struct {
    // 设备id
    DeviceId   string `json:"device_id,omitempty" xml:"device_id,omitempty"`
    // 设备名称
    DisplayName   string `json:"display_name,omitempty" xml:"display_name,omitempty"`
    // 开发商名称
    VendorName   string `json:"vendor_name,omitempty" xml:"vendor_name,omitempty"`
    // 设备类型
    DeviceType   int64 `json:"device_type,omitempty" xml:"device_type,omitempty"`
    // 是否车载设备
    IsOnVehicle   int64 `json:"is_on_vehicle,omitempty" xml:"is_on_vehicle,omitempty"`
    // 仓编码
    WarehouseCode   string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
    // 车牌号码
    PlateNumber   string `json:"plate_number,omitempty" xml:"plate_number,omitempty"`
    // 业务代码
    BusinessCode   int64 `json:"business_code,omitempty" xml:"business_code,omitempty"`
    // 安装位置
    AreaCode   int64 `json:"area_code,omitempty" xml:"area_code,omitempty"`
    // 设备分组id
    GroupId   int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
    // 分组中设备索引号
    GroupDeviceIndex   int64 `json:"group_device_index,omitempty" xml:"group_device_index,omitempty"`
}
