package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取设备状态 API请求
taobao.ailab.aicloud.top.device.getstatus

获取设备状态
*/
type TaobaoAilabAicloudTopDeviceGetstatusRequest struct {
    model.Params
    // 用户信息
    param0   *OpenBaseInfo
    // 设备id
    param1   string
}

// 初始化TaobaoAilabAicloudTopDeviceGetstatusRequest对象
func NewTaobaoAilabAicloudTopDeviceGetstatusRequest() *TaobaoAilabAicloudTopDeviceGetstatusRequest{
    return &TaobaoAilabAicloudTopDeviceGetstatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceGetstatusRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.getstatus"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceGetstatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceGetstatusRequest) SetParam0(param0 *OpenBaseInfo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r TaobaoAilabAicloudTopDeviceGetstatusRequest) GetParam0() *OpenBaseInfo {
    return r.param0
}
// Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceGetstatusRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r TaobaoAilabAicloudTopDeviceGetstatusRequest) GetParam1() string {
    return r.param1
}
