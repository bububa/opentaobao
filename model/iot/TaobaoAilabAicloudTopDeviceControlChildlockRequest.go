package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备儿童锁 APIRequest
taobao.ailab.aicloud.top.device.control.childlock

设备儿童锁
*/
type TaobaoAilabAicloudTopDeviceControlChildlockRequest struct {
    model.Params

    // 用户信息
    param0   *OpenBaseInfo 

    // 设备id
    param1   string 

    // 是否打开
    param2   bool 

}

func NewTaobaoAilabAicloudTopDeviceControlChildlockRequest() *TaobaoAilabAicloudTopDeviceControlChildlockRequest{
    return &TaobaoAilabAicloudTopDeviceControlChildlockRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopDeviceControlChildlockRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.control.childlock"
}

func (r TaobaoAilabAicloudTopDeviceControlChildlockRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopDeviceControlChildlockRequest) SetParam0(param0 *OpenBaseInfo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceControlChildlockRequest) GetParam0() *OpenBaseInfo {
    return r.param0
}

func (r *TaobaoAilabAicloudTopDeviceControlChildlockRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceControlChildlockRequest) GetParam1() string {
    return r.param1
}

func (r *TaobaoAilabAicloudTopDeviceControlChildlockRequest) SetParam2(param2 bool) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceControlChildlockRequest) GetParam2() bool {
    return r.param2
}

