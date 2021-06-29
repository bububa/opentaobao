package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
saveeventinfoforibos API请求
alibaba.campus.device.openapi.saveeventinfoforibos

IBos的事件信息上报与反馈的处理接口
*/
type AlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest struct {
    model.Params
    // 系统自动生成
    _param0   *WorkBenchContext
    // 系统自动生成
    _param1   *EventInfoApiDTO
}

// 初始化AlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest对象
func NewAlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest() *AlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest{
    return &AlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.saveeventinfoforibos"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 系统自动生成
func (r *AlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest) SetParam0(_param0 *WorkBenchContext) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest) GetParam0() *WorkBenchContext {
    return r._param0
}
// Param1 Setter
// 系统自动生成
func (r *AlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest) SetParam1(_param1 *EventInfoApiDTO) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest) GetParam1() *EventInfoApiDTO {
    return r._param1
}
