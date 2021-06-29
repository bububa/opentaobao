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
    _signName   string
    // 用户openId
    _openId   string
    // 短信模板
    _templateCode   string
    // 短信内容
    _content   string
    // 短链，替换短信内容中的${url}
    _url   string
    // 商家的APPKEY，如果openId是用商家的appKey生成的则需要传递
    _sellerAppKey   string
    // 活动或人群code
    _crowdCode   string
    // 短信拓展码
    _extendNum   string
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
func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetSignName(_signName string) error {
    r._signName = _signName
    r.Set("sign_name", _signName)
    return nil
}

// SignName Getter
func (r TaobaoJstMiniappOpenidMessageSendRequest) GetSignName() string {
    return r._signName
}
// OpenId Setter
// 用户openId
func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetOpenId(_openId string) error {
    r._openId = _openId
    r.Set("open_id", _openId)
    return nil
}

// OpenId Getter
func (r TaobaoJstMiniappOpenidMessageSendRequest) GetOpenId() string {
    return r._openId
}
// TemplateCode Setter
// 短信模板
func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetTemplateCode(_templateCode string) error {
    r._templateCode = _templateCode
    r.Set("template_code", _templateCode)
    return nil
}

// TemplateCode Getter
func (r TaobaoJstMiniappOpenidMessageSendRequest) GetTemplateCode() string {
    return r._templateCode
}
// Content Setter
// 短信内容
func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r TaobaoJstMiniappOpenidMessageSendRequest) GetContent() string {
    return r._content
}
// Url Setter
// 短链，替换短信内容中的${url}
func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetUrl(_url string) error {
    r._url = _url
    r.Set("url", _url)
    return nil
}

// Url Getter
func (r TaobaoJstMiniappOpenidMessageSendRequest) GetUrl() string {
    return r._url
}
// SellerAppKey Setter
// 商家的APPKEY，如果openId是用商家的appKey生成的则需要传递
func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetSellerAppKey(_sellerAppKey string) error {
    r._sellerAppKey = _sellerAppKey
    r.Set("seller_app_key", _sellerAppKey)
    return nil
}

// SellerAppKey Getter
func (r TaobaoJstMiniappOpenidMessageSendRequest) GetSellerAppKey() string {
    return r._sellerAppKey
}
// CrowdCode Setter
// 活动或人群code
func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetCrowdCode(_crowdCode string) error {
    r._crowdCode = _crowdCode
    r.Set("crowd_code", _crowdCode)
    return nil
}

// CrowdCode Getter
func (r TaobaoJstMiniappOpenidMessageSendRequest) GetCrowdCode() string {
    return r._crowdCode
}
// ExtendNum Setter
// 短信拓展码
func (r *TaobaoJstMiniappOpenidMessageSendRequest) SetExtendNum(_extendNum string) error {
    r._extendNum = _extendNum
    r.Set("extend_num", _extendNum)
    return nil
}

// ExtendNum Getter
func (r TaobaoJstMiniappOpenidMessageSendRequest) GetExtendNum() string {
    return r._extendNum
}
