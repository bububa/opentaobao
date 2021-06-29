package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
点播url API请求
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

// 初始化TaobaoAilabAicloudTopDeviceControlPlayurlRequest对象
func NewTaobaoAilabAicloudTopDeviceControlPlayurlRequest() *TaobaoAilabAicloudTopDeviceControlPlayurlRequest{
    return &TaobaoAilabAicloudTopDeviceControlPlayurlRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceControlPlayurlRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.control.playurl"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceControlPlayurlRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceControlPlayurlRequest) SetParam0(param0 *OpenBaseInfo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r TaobaoAilabAicloudTopDeviceControlPlayurlRequest) GetParam0() *OpenBaseInfo {
    return r.param0
}
// Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceControlPlayurlRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r TaobaoAilabAicloudTopDeviceControlPlayurlRequest) GetParam1() string {
    return r.param1
}
// Param2 Setter
// url
func (r *TaobaoAilabAicloudTopDeviceControlPlayurlRequest) SetParam2(param2 string) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

// Param2 Getter
func (r TaobaoAilabAicloudTopDeviceControlPlayurlRequest) GetParam2() string {
    return r.param2
}
