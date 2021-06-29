package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据设备uuid获取设备采集信息 APIRequest
alibaba.campus.device.openapi.getdevicerealtimelog

根据设备uuid获取设备采集信息
*/
type AlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest struct {
    model.Params

    // 请求发送端信息
    workBenchContext   *WorkBenchContext 

    // 设备uuid
    uuid   string 

}

func NewAlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest() *AlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest{
    return &AlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.getdevicerealtimelog"
}

func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest) SetWorkBenchContext(workBenchContext *WorkBenchContext) error {
    r.workBenchContext = workBenchContext
    r.Set("work_bench_context", workBenchContext)
    return nil
}

func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest) GetWorkBenchContext() *WorkBenchContext {
    return r.workBenchContext
}

func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest) GetUuid() string {
    return r.uuid
}

