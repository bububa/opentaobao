package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据设备uuid获取设备采集信息 API请求
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

// 初始化AlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest对象
func NewAlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest() *AlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest{
    return &AlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.getdevicerealtimelog"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest) SetWorkBenchContext(workBenchContext *WorkBenchContext) error {
    r.workBenchContext = workBenchContext
    r.Set("work_bench_context", workBenchContext)
    return nil
}

// WorkBenchContext Getter
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest) GetWorkBenchContext() *WorkBenchContext {
    return r.workBenchContext
}
// Uuid Setter
// 设备uuid
func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest) GetUuid() string {
    return r.uuid
}
