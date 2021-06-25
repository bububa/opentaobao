package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
点播url APIRequest
taobao.ailab.aicloud.top.device.control.playurl

点播url
*/
type TaobaoAilabAicloudTopDeviceControlPlayurlRequest struct {
    model.Params

    // 用户信息
    param0   *OpenBaseInfo 

    // 设备id
    param1   string 

    // url
    param2   string 

}

func NewTaobaoAilabAicloudTopDeviceControlPlayurlRequest() *TaobaoAilabAicloudTopDeviceControlPlayurlRequest{
    return &TaobaoAilabAicloudTopDeviceControlPlayurlRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopDeviceControlPlayurlRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.control.playurl"
}

func (r TaobaoAilabAicloudTopDeviceControlPlayurlRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopDeviceControlPlayurlRequest) SetParam0(param0 *OpenBaseInfo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceControlPlayurlRequest) GetParam0() *OpenBaseInfo {
    return r.param0
}

func (r *TaobaoAilabAicloudTopDeviceControlPlayurlRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceControlPlayurlRequest) GetParam1() string {
    return r.param1
}

func (r *TaobaoAilabAicloudTopDeviceControlPlayurlRequest) SetParam2(param2 string) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

func (r TaobaoAilabAicloudTopDeviceControlPlayurlRequest) GetParam2() string {
    return r.param2
}

