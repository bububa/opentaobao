package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
无线开放内容检查 API请求
taobao.wireless.content.check

无线开放内容检查，提供涉黄暴力政治文本检查。更详情介绍见 <a href="https://help.aliyun.com/document_detail/70439.html" target="blank">阿里云内容安全</a>
*/
type TaobaoWirelessContentCheckRequest struct {
    model.Params
    // 待检查的文本
    _text   string
}

// 初始化TaobaoWirelessContentCheckRequest对象
func NewTaobaoWirelessContentCheckRequest() *TaobaoWirelessContentCheckRequest{
    return &TaobaoWirelessContentCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWirelessContentCheckRequest) GetApiMethodName() string {
    return "taobao.wireless.content.check"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWirelessContentCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Text Setter
// 待检查的文本
func (r *TaobaoWirelessContentCheckRequest) SetText(_text string) error {
    r._text = _text
    r.Set("text", _text)
    return nil
}

// Text Getter
func (r TaobaoWirelessContentCheckRequest) GetText() string {
    return r._text
}
