package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备控制自定义扩展接口 API请求
taobao.ailab.aicloud.top.device.control.custom

设备控制自定义扩展接口
*/
type TaobaoAilabAicloudTopDeviceControlCustomAPIRequest struct {
    model.Params
    // 用户信息
    _param0   *OpenBaseInfo
    // 设备id
    _param1   string
    // 参数key-value列表
    _param2   []string
}

// 初始化TaobaoAilabAicloudTopDeviceControlCustomAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceControlCustomRequest() *TaobaoAilabAicloudTopDeviceControlCustomAPIRequest{
    return &TaobaoAilabAicloudTopDeviceControlCustomAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceControlCustomAPIRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.control.custom"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceControlCustomAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceControlCustomAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoAilabAicloudTopDeviceControlCustomAPIRequest) GetParam0() *OpenBaseInfo {
    return r._param0
}
// Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceControlCustomAPIRequest) SetParam1(_param1 string) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r TaobaoAilabAicloudTopDeviceControlCustomAPIRequest) GetParam1() string {
    return r._param1
}
// Param2 Setter
// 参数key-value列表
func (r *TaobaoAilabAicloudTopDeviceControlCustomAPIRequest) SetParam2(_param2 []string) error {
    r._param2 = _param2
    r.Set("param2", _param2)
    return nil
}

// Param2 Getter
func (r TaobaoAilabAicloudTopDeviceControlCustomAPIRequest) GetParam2() []string {
    return r._param2
}
