package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备播放暂停 APIRequest
taobao.ailab.aicloud.top.device.control.pauseandresume

设备播放暂停
*/
type TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest struct {
    model.Params

    // 用户信息
    param0   *OpenBaseInfo 

    // 设备id
    param1   string 

    // 是暂停还是继续
    param2   bool 

}

func NewTaobaoAilabAicloudTopDeviceControlPauseandresumeRequest() *TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest{
    return &TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.control.pauseandresume"
}

func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest) SetParam0(param0 *OpenBaseInfo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest) GetParam0() *OpenBaseInfo {
    return r.param0
}

func (r *TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest) GetParam1() string {
    return r.param1
}

func (r *TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest) SetParam2(param2 bool) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest) GetParam2() bool {
    return r.param2
}

