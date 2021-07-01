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
type AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest struct {
    model.Params
    // 请求发送端信息
    _workBenchContext   *WorkBenchContext
    // 设备uuid
    _uuid   string
}

// 初始化AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest对象
func NewAlibabaCampusDeviceOpenapiGetdevicerealtimelogRequest() *AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest{
    return &AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.getdevicerealtimelog"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
    r._workBenchContext = _workBenchContext
    r.Set("work_bench_context", _workBenchContext)
    return nil
}

// WorkBenchContext Getter
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) GetWorkBenchContext() *WorkBenchContext {
    return r._workBenchContext
}
// Uuid Setter
// 设备uuid
func (r *AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r AlibabaCampusDeviceOpenapiGetdevicerealtimelogAPIRequest) GetUuid() string {
    return r._uuid
}
