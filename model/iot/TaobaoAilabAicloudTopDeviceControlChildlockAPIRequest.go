package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备儿童锁 API请求
taobao.ailab.aicloud.top.device.control.childlock

设备儿童锁
*/
type TaobaoAilabAicloudTopDeviceControlChildlockAPIRequest struct {
    model.Params
    // 用户信息
    _param0   *OpenBaseInfo
    // 设备id
    _param1   string
    // 是否打开
    _param2   bool
}

// 初始化TaobaoAilabAicloudTopDeviceControlChildlockAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceControlChildlockRequest() *TaobaoAilabAicloudTopDeviceControlChildlockAPIRequest{
    return &TaobaoAilabAicloudTopDeviceControlChildlockAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceControlChildlockAPIRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.control.childlock"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceControlChildlockAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceControlChildlockAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoAilabAicloudTopDeviceControlChildlockAPIRequest) GetParam0() *OpenBaseInfo {
    return r._param0
}
// Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceControlChildlockAPIRequest) SetParam1(_param1 string) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r TaobaoAilabAicloudTopDeviceControlChildlockAPIRequest) GetParam1() string {
    return r._param1
}
// Param2 Setter
// 是否打开
func (r *TaobaoAilabAicloudTopDeviceControlChildlockAPIRequest) SetParam2(_param2 bool) error {
    r._param2 = _param2
    r.Set("param2", _param2)
    return nil
}

// Param2 Getter
func (r TaobaoAilabAicloudTopDeviceControlChildlockAPIRequest) GetParam2() bool {
    return r._param2
}
