package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单个openId用户短信发送 APIRequest
taobao.jst.miniapp.openid.message.send

单个openId用户短信发送
*/
type TaobaoJstMiniappOpenidMessageSendRequest struct {
    model.Params

    // 短信签名
    signName   string 

    // 用户openId
    openId   string 

    // 短信模板
    templateCode   string 

    // 短信内容
    content   string 

    // 短链，替换短信内容中的${url}
    url   string 

    // 商家的APPKEY，如果openId是用商家的appKey生成的则需要传递
    sellerAppKey   string 

    // 活动或人群code
    crowdCode   string 

    // 短信拓展码
    extendNum   string 

}

func NewTaobaoJstMiniappOpenidMessageSendRequest() *TaobaoJstMiniappOpenidMessageSendRequest{
    return &TaobaoJstMiniappOpenidMessageSendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstMiniappOpenidMessageSendRequest) GetApiMethodName() string {
    return "taobao.jst.miniapp.openid.message.send"
}

func (r TaobaoJstMiniappOpenidMessageSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetSignName(signName string) error {
    r.signName = signName
    r.Set("sign_name", signName)
    return nil
}

func (r TaobaoJstMiniappOpenidMessageSendRequest) GetSignName() string {
    return r.signName
}

func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetOpenId(openId string) error {
    r.openId = openId
    r.Set("open_id", openId)
    return nil
}

func (r TaobaoJstMiniappOpenidMessageSendRequest) GetOpenId() string {
    return r.openId
}

func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetTemplateCode(templateCode string) error {
    r.templateCode = templateCode
    r.Set("template_code", templateCode)
    return nil
}

func (r TaobaoJstMiniappOpenidMessageSendRequest) GetTemplateCode() string {
    return r.templateCode
}

func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r TaobaoJstMiniappOpenidMessageSendRequest) GetContent() string {
    return r.content
}

func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

func (r TaobaoJstMiniappOpenidMessageSendRequest) GetUrl() string {
    return r.url
}

func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetSellerAppKey(sellerAppKey string) error {
    r.sellerAppKey = sellerAppKey
    r.Set("seller_app_key", sellerAppKey)
    return nil
}

func (r TaobaoJstMiniappOpenidMessageSendRequest) GetSellerAppKey() string {
    return r.sellerAppKey
}

func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetCrowdCode(crowdCode string) error {
    r.crowdCode = crowdCode
    r.Set("crowd_code", crowdCode)
    return nil
}

func (r TaobaoJstMiniappOpenidMessageSendRequest) GetCrowdCode() string {
    return r.crowdCode
}

func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetExtendNum(extendNum string) error {
    r.extendNum = extendNum
    r.Set("extend_num", extendNum)
    return nil
}

func (r TaobaoJstMiniappOpenidMessageSendRequest) GetExtendNum() string {
    return r.extendNum
}

