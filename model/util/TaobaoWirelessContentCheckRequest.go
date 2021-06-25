package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
无线开放内容检查 APIRequest
taobao.wireless.content.check

无线开放内容检查，提供涉黄暴力政治文本检查。更详情介绍见 <a href="https://help.aliyun.com/document_detail/70439.html" target="blank">阿里云内容安全</a>
*/
type TaobaoWirelessContentCheckRequest struct {
    model.Params

    // 待检查的文本
    text   string 

}

func NewTaobaoWirelessContentCheckRequest() *TaobaoWirelessContentCheckRequest{
    return &TaobaoWirelessContentCheckRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWirelessContentCheckRequest) GetApiMethodName() string {
    return "taobao.wireless.content.check"
}

func (r TaobaoWirelessContentCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWirelessContentCheckRequest) SetText(text string) error {
    r.text = text
    r.Set("text", text)
    return nil
}

func (r TaobaoWirelessContentCheckRequest) GetText() string {
    return r.text
}

