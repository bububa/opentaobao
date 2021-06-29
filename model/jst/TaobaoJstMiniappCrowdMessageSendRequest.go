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
    signName   string
    // 活动code
    crowdCode   string
    // 短信模板，必须为全变量模板
    templateCode   string
    // 短信内容
    content   string
    // 短信中携带的短链，会替换短信内容中的${url}
    url   string
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
func (r *TaobaoJstMiniappCrowdMessageSendRequest) SetSignName(signName string) error {
    r.signName = signName
    r.Set("sign_name", signName)
    return nil
}

// SignName Getter
func (r TaobaoJstMiniappCrowdMessageSendRequest) GetSignName() string {
    return r.signName
}
// CrowdCode Setter
// 活动code
func (r *TaobaoJstMiniappCrowdMessageSendRequest) SetCrowdCode(crowdCode string) error {
    r.crowdCode = crowdCode
    r.Set("crowd_code", crowdCode)
    return nil
}

// CrowdCode Getter
func (r TaobaoJstMiniappCrowdMessageSendRequest) GetCrowdCode() string {
    return r.crowdCode
}
// TemplateCode Setter
// 短信模板，必须为全变量模板
func (r *TaobaoJstMiniappCrowdMessageSendRequest) SetTemplateCode(templateCode string) error {
    r.templateCode = templateCode
    r.Set("template_code", templateCode)
    return nil
}

// TemplateCode Getter
func (r TaobaoJstMiniappCrowdMessageSendRequest) GetTemplateCode() string {
    return r.templateCode
}
// Content Setter
// 短信内容
func (r *TaobaoJstMiniappCrowdMessageSendRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

// Content Getter
func (r TaobaoJstMiniappCrowdMessageSendRequest) GetContent() string {
    return r.content
}
// Url Setter
// 短信中携带的短链，会替换短信内容中的${url}
func (r *TaobaoJstMiniappCrowdMessageSendRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

// Url Getter
func (r TaobaoJstMiniappCrowdMessageSendRequest) GetUrl() string {
    return r.url
}
