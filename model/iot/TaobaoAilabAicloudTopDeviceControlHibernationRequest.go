package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定时休眠 APIRequest
taobao.ailab.aicloud.top.device.control.hibernation

定时休眠
*/
type TaobaoAilabAicloudTopDeviceControlHibernationRequest struct {
    model.Params

    // 用户信息
    param0   *OpenBaseInfo 

    // 设备id
    param1   string 

    // N秒后休眠
    param2   string 

}

func NewTaobaoAilabAicloudTopDeviceControlHibernationRequest() *TaobaoAilabAicloudTopDeviceControlHibernationRequest{
    return &TaobaoAilabAicloudTopDeviceControlHibernationRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopDeviceControlHibernationRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.control.hibernation"
}

func (r TaobaoAilabAicloudTopDeviceControlHibernationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopDeviceControlHibernationRequest) SetParam0(param0 *OpenBaseInfo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceControlHibernationRequest) GetParam0() *OpenBaseInfo {
    return r.param0
}

func (r *TaobaoAilabAicloudTopDeviceControlHibernationRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceControlHibernationRequest) GetParam1() string {
    return r.param1
}

func (r *TaobaoAilabAicloudTopDeviceControlHibernationRequest) SetParam2(param2 string) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceControlHibernationRequest) GetParam2() string {
    return r.param2
}

