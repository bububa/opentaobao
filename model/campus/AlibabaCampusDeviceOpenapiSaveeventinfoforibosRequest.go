package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
saveeventinfoforibos APIRequest
alibaba.campus.device.openapi.saveeventinfoforibos

IBos的事件信息上报与反馈的处理接口
*/
type AlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest struct {
    model.Params

    // 系统自动生成
    param0   *WorkBenchContext 

    // 系统自动生成
    param1   *EventInfoApiDto 

}

func NewAlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest() *AlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest{
    return &AlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.saveeventinfoforibos"
}

func (r AlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest) SetParam0(param0 *WorkBenchContext) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest) GetParam0() *WorkBenchContext {
    return r.param0
}

func (r *AlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest) SetParam1(param1 *EventInfoApiDto) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r AlibabaCampusDeviceOpenapiSaveeventinfoforibosRequest) GetParam1() *EventInfoApiDto {
    return r.param1
}

