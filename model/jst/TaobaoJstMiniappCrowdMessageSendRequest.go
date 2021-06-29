package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序活动短信发送 API请求
taobao.jst.miniapp.crowd.message.send

小程序活动短信发送
*/
type TaobaoJstMiniappCrowdMessageSendRequest struct {
    model.Params
    // 短信签名
    _signName   string
    // 活动code
    _crowdCode   string
    // 短信模板，必须为全变量模板
    _templateCode   string
    // 短信内容
    _content   string
    // 短信中携带的短链，会替换短信内容中的${url}
    _url   string
}

// 初始化TaobaoJstMiniappCrowdMessageSendRequest对象
func NewTaobaoJstMiniappCrowdMessageSendRequest() *TaobaoJstMiniappCrowdMessageSendRequest{
    return &TaobaoJstMiniappCrowdMessageSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstMiniappCrowdMessageSendRequest) GetApiMethodName() string {
    return "taobao.jst.miniapp.crowd.message.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstMiniappCrowdMessageSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SignName Setter
// 短信签名
func (r *TaobaoJstMiniappCrowdMessageSendRequest) SetSignName(_signName string) error {
    r._signName = _signName
    r.Set("sign_name", _signName)
    return nil
}

// SignName Getter
func (r TaobaoJstMiniappCrowdMessageSendRequest) GetSignName() string {
    return r._signName
}
// CrowdCode Setter
// 活动code
func (r *TaobaoJstMiniappCrowdMessageSendRequest) SetCrowdCode(_crowdCode string) error {
    r._crowdCode = _crowdCode
    r.Set("crowd_code", _crowdCode)
    return nil
}

// CrowdCode Getter
func (r TaobaoJstMiniappCrowdMessageSendRequest) GetCrowdCode() string {
    return r._crowdCode
}
// TemplateCode Setter
// 短信模板，必须为全变量模板
func (r *TaobaoJstMiniappCrowdMessageSendRequest) SetTemplateCode(_templateCode string) error {
    r._templateCode = _templateCode
    r.Set("template_code", _templateCode)
    return nil
}

// TemplateCode Getter
func (r TaobaoJstMiniappCrowdMessageSendRequest) GetTemplateCode() string {
    return r._templateCode
}
// Content Setter
// 短信内容
func (r *TaobaoJstMiniappCrowdMessageSendRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r TaobaoJstMiniappCrowdMessageSendRequest) GetContent() string {
    return r._content
}
// Url Setter
// 短信中携带的短链，会替换短信内容中的${url}
func (r *TaobaoJstMiniappCrowdMessageSendRequest) SetUrl(_url string) error {
    r._url = _url
    r.Set("url", _url)
    return nil
}

// Url Getter
func (r TaobaoJstMiniappCrowdMessageSendRequest) GetUrl() string {
    return r._url
}
