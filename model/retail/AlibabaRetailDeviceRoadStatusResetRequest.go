package retail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机货道解锁 API请求
alibaba.retail.device.road.status.reset

贩卖机货道解锁
*/
type AlibabaRetailDeviceRoadStatusResetRequest struct {
    model.Params
    // 设备外部编码
    deviceUuid   string
    // 阿里设备编码
    deviceCode   string
    // 阿里设备物理编码
    deviceSn   string
    // 货道编码
    roadNoList   []string
}

// 初始化AlibabaRetailDeviceRoadStatusResetRequest对象
func NewAlibabaRetailDeviceRoadStatusResetRequest() *AlibabaRetailDeviceRoadStatusResetRequest{
    return &AlibabaRetailDeviceRoadStatusResetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailDeviceRoadStatusResetRequest) GetApiMethodName() string {
    return "alibaba.retail.device.road.status.reset"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailDeviceRoadStatusResetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceUuid Setter
// 设备外部编码
func (r *AlibabaRetailDeviceRoadStatusResetRequest) SetDeviceUuid(deviceUuid string) error {
    r.deviceUuid = deviceUuid
    r.Set("device_uuid", deviceUuid)
    return nil
}

// DeviceUuid Getter
func (r AlibabaRetailDeviceRoadStatusResetRequest) GetDeviceUuid() string {
    return r.deviceUuid
}
// DeviceCode Setter
// 阿里设备编码
func (r *AlibabaRetailDeviceRoadStatusResetRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

// DeviceCode Getter
func (r AlibabaRetailDeviceRoadStatusResetRequest) GetDeviceCode() string {
    return r.deviceCode
}
// DeviceSn Setter
// 阿里设备物理编码
func (r *AlibabaRetailDeviceRoadStatusResetRequest) SetDeviceSn(deviceSn string) error {
    r.deviceSn = deviceSn
    r.Set("device_sn", deviceSn)
    return nil
}

// DeviceSn Getter
func (r AlibabaRetailDeviceRoadStatusResetRequest) GetDeviceSn() string {
    return r.deviceSn
}
// RoadNoList Setter
// 货道编码
func (r *AlibabaRetailDeviceRoadStatusResetRequest) SetRoadNoList(roadNoList []string) error {
    r.roadNoList = roadNoList
    r.Set("road_no_list", roadNoList)
    return nil
}

// RoadNoList Getter
func (r AlibabaRetailDeviceRoadStatusResetRequest) GetRoadNoList() []string {
    return r.roadNoList
}
