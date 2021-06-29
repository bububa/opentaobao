package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单个openId用户短信发送 API请求
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

// 初始化TaobaoJstMiniappOpenidMessageSendRequest对象
func NewTaobaoJstMiniappOpenidMessageSendRequest() *TaobaoJstMiniappOpenidMessageSendRequest{
    return &TaobaoJstMiniappOpenidMessageSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstMiniappOpenidMessageSendRequest) GetApiMethodName() string {
    return "taobao.jst.miniapp.openid.message.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstMiniappOpenidMessageSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SignName Setter
// 短信签名
func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetSignName(signName string) error {
    r.signName = signName
    r.Set("sign_name", signName)
    return nil
}

// SignName Getter
func (r TaobaoJstMiniappOpenidMessageSendRequest) GetSignName() string {
    return r.signName
}
// OpenId Setter
// 用户openId
func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetOpenId(openId string) error {
    r.openId = openId
    r.Set("open_id", openId)
    return nil
}

// OpenId Getter
func (r TaobaoJstMiniappOpenidMessageSendRequest) GetOpenId() string {
    return r.openId
}
// TemplateCode Setter
// 短信模板
func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetTemplateCode(templateCode string) error {
    r.templateCode = templateCode
    r.Set("template_code", templateCode)
    return nil
}

// TemplateCode Getter
func (r TaobaoJstMiniappOpenidMessageSendRequest) GetTemplateCode() string {
    return r.templateCode
}
// Content Setter
// 短信内容
func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

// Content Getter
func (r TaobaoJstMiniappOpenidMessageSendRequest) GetContent() string {
    return r.content
}
// Url Setter
// 短链，替换短信内容中的${url}
func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

// Url Getter
func (r TaobaoJstMiniappOpenidMessageSendRequest) GetUrl() string {
    return r.url
}
// SellerAppKey Setter
// 商家的APPKEY，如果openId是用商家的appKey生成的则需要传递
func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetSellerAppKey(sellerAppKey string) error {
    r.sellerAppKey = sellerAppKey
    r.Set("seller_app_key", sellerAppKey)
    return nil
}

// SellerAppKey Getter
func (r TaobaoJstMiniappOpenidMessageSendRequest) GetSellerAppKey() string {
    return r.sellerAppKey
}
// CrowdCode Setter
// 活动或人群code
func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetCrowdCode(crowdCode string) error {
    r.crowdCode = crowdCode
    r.Set("crowd_code", crowdCode)
    return nil
}

// CrowdCode Getter
func (r TaobaoJstMiniappOpenidMessageSendRequest) GetCrowdCode() string {
    return r.crowdCode
}
// ExtendNum Setter
// 短信拓展码
func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetExtendNum(extendNum string) error {
    r.extendNum = extendNum
    r.Set("extend_num", extendNum)
    return nil
}

// ExtendNum Getter
func (r TaobaoJstMiniappOpenidMessageSendRequest) GetExtendNum() string {
    return r.extendNum
}
