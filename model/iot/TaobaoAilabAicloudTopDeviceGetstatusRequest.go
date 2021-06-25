package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备状态 APIRequest
taobao.ailab.aicloud.top.device.getstatus

获取设备状态
*/
type TaobaoAilabAicloudTopDeviceGetstatusRequest struct {
    model.Params

    // 用户信息
    param0   *OpenBaseInfo 

    // 设备id
    param1   string 

}

func NewTaobaoAilabAicloudTopDeviceGetstatusRequest() *TaobaoAilabAicloudTopDeviceGetstatusRequest{
    return &TaobaoAilabAicloudTopDeviceGetstatusRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopDeviceGetstatusRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.getstatus"
}

func (r TaobaoAilabAicloudTopDeviceGetstatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopDeviceGetstatusRequest) SetParam0(param0 *OpenBaseInfo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceGetstatusRequest) GetParam0() *OpenBaseInfo {
    return r.param0
}

func (r *TaobaoAilabAicloudTopDeviceGetstatusRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceGetstatusRequest) GetParam1() string {
    return r.param1
}

