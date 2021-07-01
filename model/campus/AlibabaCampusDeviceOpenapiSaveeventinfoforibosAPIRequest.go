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
type AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest struct {
    model.Params
    // 系统自动生成
    _param0   *WorkBenchContext
    // 系统自动生成
    _param1   *EventInfoApiDto
}

// 初始化AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest对象
func NewAlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest() *AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest{
    return &AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.saveeventinfoforibos"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 系统自动生成
func (r *AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest) GetParam0() *WorkBenchContext {
    return r._param0
}
// Param1 Setter
// 系统自动生成
func (r *AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest) SetParam1(_param1 *EventInfoApiDto) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r AlibabaCampusDeviceOpenapiSaveeventinfoforibosAPIRequest) GetParam1() *EventInfoApiDto {
    return r._param1
}
