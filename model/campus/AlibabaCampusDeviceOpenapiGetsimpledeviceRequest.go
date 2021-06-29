package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单个设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息) APIRequest
alibaba.campus.device.openapi.getsimpledevice

获取指定设备的基础信息(仅包含设备id,code,是否启用,位置信息,描述等基础信息)
*/
type AlibabaCampusDeviceOpenapiGetsimpledeviceRequest struct {
    model.Params

    // 请求发送端信息
    workBenchContext   *WorkBenchContext 

    // 设备uuid
    uuid   string 

}

func NewAlibabaCampusDeviceOpenapiGetsimpledeviceRequest() *AlibabaCampusDeviceOpenapiGetsimpledeviceRequest{
    return &AlibabaCampusDeviceOpenapiGetsimpledeviceRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusDeviceOpenapiGetsimpledeviceRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.getsimpledevice"
}

func (r AlibabaCampusDeviceOpenapiGetsimpledeviceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusDeviceOpenapiGetsimpledeviceRequest) SetWorkBenchContext(workBenchContext *WorkBenchContext) error {
    r.workBenchContext = workBenchContext
    r.Set("work_bench_context", workBenchContext)
    return nil
}

func (r AlibabaCampusDeviceOpenapiGetsimpledeviceRequest) GetWorkBenchContext() *WorkBenchContext {
    return r.workBenchContext
}

func (r *AlibabaCampusDeviceOpenapiGetsimpledeviceRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r AlibabaCampusDeviceOpenapiGetsimpledeviceRequest) GetUuid() string {
    return r.uuid
}

