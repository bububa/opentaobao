package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备音量 API请求
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

// 初始化TaobaoAilabAicloudTopDeviceControlVolumeRequest对象
func NewTaobaoAilabAicloudTopDeviceControlVolumeRequest() *TaobaoAilabAicloudTopDeviceControlVolumeRequest{
    return &TaobaoAilabAicloudTopDeviceControlVolumeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceControlVolumeRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.control.volume"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceControlVolumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceControlVolumeRequest) SetParam0(param0 *OpenBaseInfo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r TaobaoAilabAicloudTopDeviceControlVolumeRequest) GetParam0() *OpenBaseInfo {
    return r.param0
}
// Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceControlVolumeRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r TaobaoAilabAicloudTopDeviceControlVolumeRequest) GetParam1() string {
    return r.param1
}
// Param2 Setter
// 音量0-100
func (r *TaobaoAilabAicloudTopDeviceControlVolumeRequest) SetParam2(param2 int64) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

// Param2 Getter
func (r TaobaoAilabAicloudTopDeviceControlVolumeRequest) GetParam2() int64 {
    return r.param2
}
