package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定时休眠 API请求
taobao.ailab.aicloud.top.device.control.hibernation

定时休眠
*/
type TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest struct {
    model.Params
    // 用户信息
    _param0   *OpenBaseInfo
    // 设备id
    _param1   string
    // N秒后休眠
    _param2   string
}

// 初始化TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceControlHibernationRequest() *TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest{
    return &TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.control.hibernation"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest) GetParam0() *OpenBaseInfo {
    return r._param0
}
// Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest) SetParam1(_param1 string) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest) GetParam1() string {
    return r._param1
}
// Param2 Setter
// N秒后休眠
func (r *TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest) SetParam2(_param2 string) error {
    r._param2 = _param2
    r.Set("param2", _param2)
    return nil
}

// Param2 Getter
func (r TaobaoAilabAicloudTopDeviceControlHibernationAPIRequest) GetParam2() string {
    return r._param2
}
