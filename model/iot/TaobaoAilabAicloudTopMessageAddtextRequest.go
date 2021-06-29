package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
精灵代说 API请求
taobao.ailab.aicloud.top.message.addtext

精灵代说
*/
type TaobaoAilabAicloudTopMessageAddtextRequest struct {
    model.Params
    // 用户信息
    param0   *OpenBaseInfo
    // 设备id
    param1   string
    // 代说文本
    param2   string
    // 扩展信息，可以为空
    param3   string
}

// 初始化TaobaoAilabAicloudTopMessageAddtextRequest对象
func NewTaobaoAilabAicloudTopMessageAddtextRequest() *TaobaoAilabAicloudTopMessageAddtextRequest{
    return &TaobaoAilabAicloudTopMessageAddtextRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMessageAddtextRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.message.addtext"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMessageAddtextRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 用户信息
func (r *TaobaoAilabAicloudTopMessageAddtextRequest) SetParam0(param0 *OpenBaseInfo) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r TaobaoAilabAicloudTopMessageAddtextRequest) GetParam0() *OpenBaseInfo {
    return r.param0
}
// Param1 Setter
// 设备id
func (r *TaobaoAilabAicloudTopMessageAddtextRequest) SetParam1(param1 string) error {
    r.param1 = param1
    r.Set("param1", param1)
    return nil
}

// Param1 Getter
func (r TaobaoAilabAicloudTopMessageAddtextRequest) GetParam1() string {
    return r.param1
}
// Param2 Setter
// 代说文本
func (r *TaobaoAilabAicloudTopMessageAddtextRequest) SetParam2(param2 string) error {
    r.param2 = param2
    r.Set("param2", param2)
    return nil
}

// Param2 Getter
func (r TaobaoAilabAicloudTopMessageAddtextRequest) GetParam2() string {
    return r.param2
}
// Param3 Setter
// 扩展信息，可以为空
func (r *TaobaoAilabAicloudTopMessageAddtextRequest) SetParam3(param3 string) error {
    r.param3 = param3
    r.Set("param3", param3)
    return nil
}

// Param3 Getter
func (r TaobaoAilabAicloudTopMessageAddtextRequest) GetParam3() string {
    return r.param3
}
