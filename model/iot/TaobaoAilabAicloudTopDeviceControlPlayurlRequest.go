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
type TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest struct {
    model.Params
    // 用户信息
    _param0   *OpenBaseInfo
    // 设备id
    _param1   string
    // url
    _param2   string
}

// 初始化TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceControlPlayurlRequest() *TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest{
    return &TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.control.playurl"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest) GetParam0() *OpenBaseInfo {
    return r._param0
}
// Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest) SetParam1(_param1 string) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest) GetParam1() string {
    return r._param1
}
// Param2 Setter
// url
func (r *TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest) SetParam2(_param2 string) error {
    r._param2 = _param2
    r.Set("param2", _param2)
    return nil
}

// Param2 Getter
func (r TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest) GetParam2() string {
    return r._param2
}
