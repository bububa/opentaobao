package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过id播放歌曲 API请求
taobao.ailab.aicloud.top.device.control.playbyid

通过id播放歌曲
*/
type TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest struct {
    model.Params
    // 用户信息
    _param0   *OpenBaseInfo
    // 设备id
    _param1   string
    // 音频id
    _param2   string
    // 音频来源
    _param3   string
    // 音频类型，如果没有音频类型默认填children_song
    _param4   string
}

// 初始化TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceControlPlaybyidRequest() *TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest{
    return &TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.device.control.playbyid"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest) SetParam0(_param0 *OpenBaseInfo) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest) GetParam0() *OpenBaseInfo {
    return r._param0
}
// Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest) SetParam1(_param1 string) error {
    r._param1 = _param1
    r.Set("param1", _param1)
    return nil
}

// Param1 Getter
func (r TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest) GetParam1() string {
    return r._param1
}
// Param2 Setter
// 音频id
func (r *TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest) SetParam2(_param2 string) error {
    r._param2 = _param2
    r.Set("param2", _param2)
    return nil
}

// Param2 Getter
func (r TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest) GetParam2() string {
    return r._param2
}
// Param3 Setter
// 音频来源
func (r *TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest) SetParam3(_param3 string) error {
    r._param3 = _param3
    r.Set("param3", _param3)
    return nil
}

// Param3 Getter
func (r TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest) GetParam3() string {
    return r._param3
}
// Param4 Setter
// 音频类型，如果没有音频类型默认填children_song
func (r *TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest) SetParam4(_param4 string) error {
    r._param4 = _param4
    r.Set("param4", _param4)
    return nil
}

// Param4 Getter
func (r TaobaoAilabAicloudTopDeviceControlPlaybyidAPIRequest) GetParam4() string {
    return r._param4
}
