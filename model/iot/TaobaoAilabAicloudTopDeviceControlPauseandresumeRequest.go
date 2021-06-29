package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设备播放暂停 API请求
taobao.ailab.aicloud.top.device.control.pauseandresume

设备播放暂停
*/
type TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest struct {
    model.Params
    // 用户信息
    param0   *OpenBaseInfo
    // 设备id
    param1   string
    // 是暂停还是继续
    param2   bool
}

// 初始化TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest对象
func NewTaobaoAilabAicloudTopDeviceControlPauseandresumeRequest() *TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest{
    return &TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.control.pauseandresume"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest) SetParam0(param0 *OpenBaseInfo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest) GetParam0() *OpenBaseInfo {
    return r.param0
}
// Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest) GetParam1() string {
    return r.param1
}
// Param2 Setter
// 是暂停还是继续
func (r *TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest) SetParam2(param2 bool) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

// Param2 Getter
func (r TaobaoAilabAicloudTopDeviceControlPauseandresumeRequest) GetParam2() bool {
    return r.param2
}
