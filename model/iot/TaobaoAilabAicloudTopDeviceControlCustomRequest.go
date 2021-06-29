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
type TaobaoAilabAicloudTopDeviceControlCustomRequest struct {
    model.Params
    // 用户信息
    param0   *OpenBaseInfo
    // 设备id
    param1   string
    // 参数key-value列表
    param2   []string
}

// 初始化TaobaoAilabAicloudTopDeviceControlCustomRequest对象
func NewTaobaoAilabAicloudTopDeviceControlCustomRequest() *TaobaoAilabAicloudTopDeviceControlCustomRequest{
    return &TaobaoAilabAicloudTopDeviceControlCustomRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceControlCustomRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.control.custom"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceControlCustomRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceControlCustomRequest) SetParam0(param0 *OpenBaseInfo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r TaobaoAilabAicloudTopDeviceControlCustomRequest) GetParam0() *OpenBaseInfo {
    return r.param0
}
// Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceControlCustomRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r TaobaoAilabAicloudTopDeviceControlCustomRequest) GetParam1() string {
    return r.param1
}
// Param2 Setter
// 参数key-value列表
func (r *TaobaoAilabAicloudTopDeviceControlCustomRequest) SetParam2(param2 []string) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

// Param2 Getter
func (r TaobaoAilabAicloudTopDeviceControlCustomRequest) GetParam2() []string {
    return r.param2
}
