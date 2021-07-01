package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备数据上报 API请求
alibaba.campus.devicehub.openapi.reportdata

设备数据上报
*/
type AlibabaCampusDevicehubOpenapiReportdataAPIRequest struct {
    model.Params
    // 自动生成
    _deviceEventData   *DeviceReportEventDTO
}

// 初始化AlibabaCampusDevicehubOpenapiReportdataAPIRequest对象
func NewAlibabaCampusDevicehubOpenapiReportdataRequest() *AlibabaCampusDevicehubOpenapiReportdataAPIRequest{
    return &AlibabaCampusDevicehubOpenapiReportdataAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusDevicehubOpenapiReportdataAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.devicehub.openapi.reportdata"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusDevicehubOpenapiReportdataAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceEventData Setter
// 自动生成
func (r *AlibabaCampusDevicehubOpenapiReportdataAPIRequest) SetDeviceEventData(_deviceEventData *DeviceReportEventDTO) error {
    r._deviceEventData = _deviceEventData
    r.Set("device_event_data", _deviceEventData)
    return nil
}

// DeviceEventData Getter
func (r AlibabaCampusDevicehubOpenapiReportdataAPIRequest) GetDeviceEventData() *DeviceReportEventDTO {
    return r._deviceEventData
}
