package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据设备uuid获取设备信息 APIRequest
alibaba.campus.device.openapi.getuniquedevice

根据设备uuid获取设备信息
*/
type AlibabaCampusDeviceOpenapiGetuniquedeviceRequest struct {
    model.Params

    // 请求发送端信息
    workBenchContext   *WorkBenchContext 

    // 设备序列号uuid
    uuid   string 

}

func NewAlibabaCampusDeviceOpenapiGetuniquedeviceRequest() *AlibabaCampusDeviceOpenapiGetuniquedeviceRequest{
    return &AlibabaCampusDeviceOpenapiGetuniquedeviceRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusDeviceOpenapiGetuniquedeviceRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.getuniquedevice"
}

func (r AlibabaCampusDeviceOpenapiGetuniquedeviceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusDeviceOpenapiGetuniquedeviceRequest) SetWorkBenchContext(workBenchContext *WorkBenchContext) error {
    r.workBenchContext = workBenchContext
    r.Set("work_bench_context", workBenchContext)
    return nil
}

func (r AlibabaCampusDeviceOpenapiGetuniquedeviceRequest) GetWorkBenchContext() *WorkBenchContext {
    return r.workBenchContext
}

func (r *AlibabaCampusDeviceOpenapiGetuniquedeviceRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r AlibabaCampusDeviceOpenapiGetuniquedeviceRequest) GetUuid() string {
    return r.uuid
}

