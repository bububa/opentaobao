package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
台灯控制 APIRequest
taobao.ailab.aicloud.top.device.control.lamp

台灯控制
*/
type TaobaoAilabAicloudTopDeviceControlLampRequest struct {
    model.Params

    // 用户信息
    param0   *OpenBaseInfo 

    // 设备id
    param1   string 

    // 是否打开
    param2   bool 

    // 目标名称
    param3   string 

}

func NewTaobaoAilabAicloudTopDeviceControlLampRequest() *TaobaoAilabAicloudTopDeviceControlLampRequest{
    return &TaobaoAilabAicloudTopDeviceControlLampRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopDeviceControlLampRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.control.lamp"
}

func (r TaobaoAilabAicloudTopDeviceControlLampRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopDeviceControlLampRequest) SetParam0(param0 *OpenBaseInfo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceControlLampRequest) GetParam0() *OpenBaseInfo {
    return r.param0
}

func (r *TaobaoAilabAicloudTopDeviceControlLampRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceControlLampRequest) GetParam1() string {
    return r.param1
}

func (r *TaobaoAilabAicloudTopDeviceControlLampRequest) SetParam2(param2 bool) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceControlLampRequest) GetParam2() bool {
    return r.param2
}

func (r *TaobaoAilabAicloudTopDeviceControlLampRequest) SetParam3(param3 string) error {
    r.param3 = param3
    r.Set("param3", param3)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceControlLampRequest) GetParam3() string {
    return r.param3
}

