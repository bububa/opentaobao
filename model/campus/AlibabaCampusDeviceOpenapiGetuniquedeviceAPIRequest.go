package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据设备uuid获取设备信息 API请求
alibaba.campus.device.openapi.getuniquedevice

根据设备uuid获取设备信息
*/
type AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest struct {
    model.Params
    // 请求发送端信息
    _workBenchContext   *WorkBenchContext
    // 设备序列号uuid
    _uuid   string
}

// 初始化AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest对象
func NewAlibabaCampusDeviceOpenapiGetuniquedeviceRequest() *AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest{
    return &AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.getuniquedevice"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkBenchContext Setter
// 请求发送端信息
func (r *AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
    r._workBenchContext = _workBenchContext
    r.Set("work_bench_context", _workBenchContext)
    return nil
}

// WorkBenchContext Getter
func (r AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) GetWorkBenchContext() *WorkBenchContext {
    return r._workBenchContext
}
// Uuid Setter
// 设备序列号uuid
func (r *AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r AlibabaCampusDeviceOpenapiGetuniquedeviceAPIRequest) GetUuid() string {
    return r._uuid
}
