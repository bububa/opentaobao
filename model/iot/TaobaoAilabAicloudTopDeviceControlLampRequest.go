package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
台灯控制 API请求
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

// 初始化TaobaoAilabAicloudTopDeviceControlLampRequest对象
func NewTaobaoAilabAicloudTopDeviceControlLampRequest() *TaobaoAilabAicloudTopDeviceControlLampRequest{
    return &TaobaoAilabAicloudTopDeviceControlLampRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceControlLampRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.control.lamp"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceControlLampRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceControlLampRequest) SetParam0(param0 *OpenBaseInfo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r TaobaoAilabAicloudTopDeviceControlLampRequest) GetParam0() *OpenBaseInfo {
    return r.param0
}
// Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceControlLampRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r TaobaoAilabAicloudTopDeviceControlLampRequest) GetParam1() string {
    return r.param1
}
// Param2 Setter
// 是否打开
func (r *TaobaoAilabAicloudTopDeviceControlLampRequest) SetParam2(param2 bool) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

// Param2 Getter
func (r TaobaoAilabAicloudTopDeviceControlLampRequest) GetParam2() bool {
    return r.param2
}
// Param3 Setter
// 目标名称
func (r *TaobaoAilabAicloudTopDeviceControlLampRequest) SetParam3(param3 string) error {
    r.param3 = param3
    r.Set("param3", param3)
    return nil
}

// Param3 Getter
func (r TaobaoAilabAicloudTopDeviceControlLampRequest) GetParam3() string {
    return r.param3
}
