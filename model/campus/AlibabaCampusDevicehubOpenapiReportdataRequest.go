package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备数据上报 APIRequest
alibaba.campus.devicehub.openapi.reportdata

设备数据上报
*/
type AlibabaCampusDevicehubOpenapiReportdataRequest struct {
    model.Params

    // 自动生成
    deviceEventData   *DeviceReportEventDTO 

}

func NewAlibabaCampusDevicehubOpenapiReportdataRequest() *AlibabaCampusDevicehubOpenapiReportdataRequest{
    return &AlibabaCampusDevicehubOpenapiReportdataRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusDevicehubOpenapiReportdataRequest) GetApiMethodName() string {
    return "alibaba.campus.devicehub.openapi.reportdata"
}

func (r AlibabaCampusDevicehubOpenapiReportdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusDevicehubOpenapiReportdataRequest) SetDeviceEventData(deviceEventData *DeviceReportEventDTO) error {
    r.deviceEventData = deviceEventData
    r.Set("device_event_data", deviceEventData)
    return nil
}

func (r AlibabaCampusDevicehubOpenapiReportdataRequest) GetDeviceEventData() *DeviceReportEventDTO {
    return r.deviceEventData
}

