package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstMiniappOpenidMessageSendAPIRequest 单个openId用户短信发送 API请求
// taobao.jst.miniapp.openid.message.send
//
// 单个openId用户短信发送
type TaobaoJstMiniappOpenidMessageSendAPIRequest struct {
	model.Params
	// 短信签名
	_signName string
	// 用户openId
	_openId string
	// 短信模板
	_templateCode string
	// 短信内容
	_content string
	// 短链，替换短信内容中的${url}
	_url string
	// 商家的APPKEY，如果openId是用商家的appKey生成的则需要传递
	_sellerAppKey string
	// 活动或人群code
	_crowdCode string
	// 短信拓展码
	_extendNum string
}

// NewTaobaoJstMiniappOpenidMessageSendRequest 初始化TaobaoJstMiniappOpenidMessageSendAPIRequest对象
func NewTaobaoJstMiniappOpenidMessageSendRequest() *TaobaoJstMiniappOpenidMessageSendAPIRequest {
	return &TaobaoJstMiniappOpenidMessageSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstMiniappOpenidMessageSendAPIRequest) GetApiMethodName() string {
	return "taobao.jst.miniapp.openid.message.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstMiniappOpenidMessageSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstMiniappOpenidMessageSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSignName is SignName Setter
// 短信签名
func (r *TaobaoJstMiniappOpenidMessageSendAPIRequest) SetSignName(_signName string) error {
	r._signName = _signName
	r.Set("sign_name", _signName)
	return nil
}

// GetSignName SignName Getter
func (r TaobaoJstMiniappOpenidMessageSendAPIRequest) GetSignName() string {
	return r._signName
}

// SetOpenId is OpenId Setter
// 用户openId
func (r *TaobaoJstMiniappOpenidMessageSendAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r TaobaoJstMiniappOpenidMessageSendAPIRequest) GetOpenId() string {
	return r._openId
}

// SetTemplateCode is TemplateCode Setter
// 短信模板
func (r *TaobaoJstMiniappOpenidMessageSendAPIRequest) SetTemplateCode(_templateCode string) error {
	r._templateCode = _templateCode
	r.Set("template_code", _templateCode)
	return nil
}

// GetTemplateCode TemplateCode Getter
func (r TaobaoJstMiniappOpenidMessageSendAPIRequest) GetTemplateCode() string {
	return r._templateCode
}

// SetContent is Content Setter
// 短信内容
func (r *TaobaoJstMiniappOpenidMessageSendAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaoJstMiniappOpenidMessageSendAPIRequest) GetContent() string {
	return r._content
}

// SetUrl is Url Setter
// 短链，替换短信内容中的${url}
func (r *TaobaoJstMiniappOpenidMessageSendAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// GetUrl Url Getter
func (r TaobaoJstMiniappOpenidMessageSendAPIRequest) GetUrl() string {
	return r._url
}

// SetSellerAppKey is SellerAppKey Setter
// 商家的APPKEY，如果openId是用商家的appKey生成的则需要传递
func (r *TaobaoJstMiniappOpenidMessageSendAPIRequest) SetSellerAppKey(_sellerAppKey string) error {
	r._sellerAppKey = _sellerAppKey
	r.Set("seller_app_key", _sellerAppKey)
	return nil
}

// GetSellerAppKey SellerAppKey Getter
func (r TaobaoJstMiniappOpenidMessageSendAPIRequest) GetSellerAppKey() string {
	return r._sellerAppKey
}

// SetCrowdCode is CrowdCode Setter
// 活动或人群code
func (r *TaobaoJstMiniappOpenidMessageSendAPIRequest) SetCrowdCode(_crowdCode string) error {
	r._crowdCode = _crowdCode
	r.Set("crowd_code", _crowdCode)
	return nil
}

// GetCrowdCode CrowdCode Getter
func (r TaobaoJstMiniappOpenidMessageSendAPIRequest) GetCrowdCode() string {
	return r._crowdCode
}

// SetExtendNum is ExtendNum Setter
// 短信拓展码
func (r *TaobaoJstMiniappOpenidMessageSendAPIRequest) SetExtendNum(_extendNum string) error {
	r._extendNum = _extendNum
	r.Set("extend_num", _extendNum)
	return nil
}

// GetExtendNum ExtendNum Getter
func (r TaobaoJstMiniappOpenidMessageSendAPIRequest) GetExtendNum() string {
	return r._extendNum
}
