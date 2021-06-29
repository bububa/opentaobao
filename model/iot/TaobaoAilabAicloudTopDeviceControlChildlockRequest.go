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
type TaobaoAilabAicloudTopDeviceControlChildlockRequest struct {
    model.Params
    // 用户信息
    _param0   *OpenBaseInfo
    // 设备id
    _param1   string
    // 是否打开
    _param2   bool
}

// 初始化TaobaoAilabAicloudTopDeviceControlChildlockRequest对象
func NewTaobaoAilabAicloudTopDeviceControlChildlockRequest() *TaobaoAilabAicloudTopDeviceControlChildlockRequest{
    return &TaobaoAilabAicloudTopDeviceControlChildlockRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceControlChildlockRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.control.childlock"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceControlChildlockRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceControlChildlockRequest) SetParam0(_param0 *OpenBaseInfo) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoAilabAicloudTopDeviceControlChildlockRequest) GetParam0() *OpenBaseInfo {
    return r._param0
}
// Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceControlChildlockRequest) SetParam1(_param1 string) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r TaobaoAilabAicloudTopDeviceControlChildlockRequest) GetParam1() string {
    return r._param1
}
// Param2 Setter
// 是否打开
func (r *TaobaoAilabAicloudTopDeviceControlChildlockRequest) SetParam2(_param2 bool) error {
    r._param2 = _param2
    r.Set("param2", _param2)
    return nil
}

// Param2 Getter
func (r TaobaoAilabAicloudTopDeviceControlChildlockRequest) GetParam2() bool {
    return r._param2
}
