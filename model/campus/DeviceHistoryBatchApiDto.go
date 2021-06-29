package campus

// DeviceHistoryBatchApiDto 
type DeviceHistoryBatchApiDto struct {
    // 设备参数code
    PropertyCode   string `json:"property_code,omitempty" xml:"property_code,omitempty"`
    // 设备id
    DeviceId   string `json:"device_id,omitempty" xml:"device_id,omitempty"`
    // 当前设备查询结果状态
    Success   string `json:"success,omitempty" xml:"success,omitempty"`
    // 当前查询错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 当前查询错误编码
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // 当前设备运行参数时间切片历史数据
    DeviceIntervalValue   string `json:"device_interval_value,omitempty" xml:"device_interval_value,omitempty"`
}
