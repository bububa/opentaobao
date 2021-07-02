package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstMiniappCrowdMessageSendAPIRequest 小程序活动短信发送 API请求
// taobao.jst.miniapp.crowd.message.send
//
// 小程序活动短信发送
type TaobaoJstMiniappCrowdMessageSendAPIRequest struct {
	model.Params
	// 短信签名
	_signName string
	// 活动code
	_crowdCode string
	// 短信模板，必须为全变量模板
	_templateCode string
	// 短信内容
	_content string
	// 短信中携带的短链，会替换短信内容中的${url}
	_url string
}

// NewTaobaoJstMiniappCrowdMessageSendRequest 初始化TaobaoJstMiniappCrowdMessageSendAPIRequest对象
func NewTaobaoJstMiniappCrowdMessageSendRequest() *TaobaoJstMiniappCrowdMessageSendAPIRequest {
	return &TaobaoJstMiniappCrowdMessageSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstMiniappCrowdMessageSendAPIRequest) GetApiMethodName() string {
	return "taobao.jst.miniapp.crowd.message.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstMiniappCrowdMessageSendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSignName is SignName Setter
// 短信签名
func (r *TaobaoJstMiniappCrowdMessageSendAPIRequest) SetSignName(_signName string) error {
	r._signName = _signName
	r.Set("sign_name", _signName)
	return nil
}

// GetSignName SignName Getter
func (r TaobaoJstMiniappCrowdMessageSendAPIRequest) GetSignName() string {
	return r._signName
}

// SetCrowdCode is CrowdCode Setter
// 活动code
func (r *TaobaoJstMiniappCrowdMessageSendAPIRequest) SetCrowdCode(_crowdCode string) error {
	r._crowdCode = _crowdCode
	r.Set("crowd_code", _crowdCode)
	return nil
}

// GetCrowdCode CrowdCode Getter
func (r TaobaoJstMiniappCrowdMessageSendAPIRequest) GetCrowdCode() string {
	return r._crowdCode
}

// SetTemplateCode is TemplateCode Setter
// 短信模板，必须为全变量模板
func (r *TaobaoJstMiniappCrowdMessageSendAPIRequest) SetTemplateCode(_templateCode string) error {
	r._templateCode = _templateCode
	r.Set("template_code", _templateCode)
	return nil
}

// GetTemplateCode TemplateCode Getter
func (r TaobaoJstMiniappCrowdMessageSendAPIRequest) GetTemplateCode() string {
	return r._templateCode
}

// SetContent is Content Setter
// 短信内容
func (r *TaobaoJstMiniappCrowdMessageSendAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaoJstMiniappCrowdMessageSendAPIRequest) GetContent() string {
	return r._content
}

// SetUrl is Url Setter
// 短信中携带的短链，会替换短信内容中的${url}
func (r *TaobaoJstMiniappCrowdMessageSendAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// GetUrl Url Getter
func (r TaobaoJstMiniappCrowdMessageSendAPIRequest) GetUrl() string {
	return r._url
}
