package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序活动短信发送 APIRequest
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

func NewTaobaoJstMiniappCrowdMessageSendRequest() *TaobaoJstMiniappCrowdMessageSendRequest{
    return &TaobaoJstMiniappCrowdMessageSendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstMiniappCrowdMessageSendRequest) GetApiMethodName() string {
    return "taobao.jst.miniapp.crowd.message.send"
}

func (r TaobaoJstMiniappCrowdMessageSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstMiniappCrowdMessageSendRequest) SetSignName(signName string) error {
    r.signName = signName
    r.Set("sign_name", signName)
    return nil
}

func (r TaobaoJstMiniappCrowdMessageSendRequest) GetSignName() string {
    return r.signName
}

func (r *TaobaoJstMiniappCrowdMessageSendRequest) SetCrowdCode(crowdCode string) error {
    r.crowdCode = crowdCode
    r.Set("crowd_code", crowdCode)
    return nil
}

func (r TaobaoJstMiniappCrowdMessageSendRequest) GetCrowdCode() string {
    return r.crowdCode
}

func (r *TaobaoJstMiniappCrowdMessageSendRequest) SetTemplateCode(templateCode string) error {
    r.templateCode = templateCode
    r.Set("template_code", templateCode)
    return nil
}

func (r TaobaoJstMiniappCrowdMessageSendRequest) GetTemplateCode() string {
    return r.templateCode
}

func (r *TaobaoJstMiniappCrowdMessageSendRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r TaobaoJstMiniappCrowdMessageSendRequest) GetContent() string {
    return r.content
}

func (r *TaobaoJstMiniappCrowdMessageSendRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

func (r TaobaoJstMiniappCrowdMessageSendRequest) GetUrl() string {
    return r.url
}

