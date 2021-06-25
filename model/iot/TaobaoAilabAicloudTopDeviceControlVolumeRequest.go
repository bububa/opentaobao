package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备音量 APIRequest
taobao.ailab.aicloud.top.device.control.volume

设备音量
*/
type TaobaoAilabAicloudTopDeviceControlVolumeRequest struct {
    model.Params

    // 用户信息
    param0   *OpenBaseInfo 

    // 设备id
    param1   string 

    // 音量0-100
    param2   int64 

}

func NewTaobaoAilabAicloudTopDeviceControlVolumeRequest() *TaobaoAilabAicloudTopDeviceControlVolumeRequest{
    return &TaobaoAilabAicloudTopDeviceControlVolumeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopDeviceControlVolumeRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.control.volume"
}

func (r TaobaoAilabAicloudTopDeviceControlVolumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopDeviceControlVolumeRequest) SetParam0(param0 *OpenBaseInfo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceControlVolumeRequest) GetParam0() *OpenBaseInfo {
    return r.param0
}

func (r *TaobaoAilabAicloudTopDeviceControlVolumeRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceControlVolumeRequest) GetParam1() string {
    return r.param1
}

func (r *TaobaoAilabAicloudTopDeviceControlVolumeRequest) SetParam2(param2 int64) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceControlVolumeRequest) GetParam2() int64 {
    return r.param2
}

