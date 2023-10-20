package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportAuthUnbindAPIRequest 三方关系解绑接口 API请求
// alibaba.alisports.passport.auth.unbind
//
// 解除阿里体育openId和三方合作方的openId的绑定关系
type AlibabaAlisportsPassportAuthUnbindAPIRequest struct {
	model.Params
	// 阿里体育业务KEY
	_alispAppKey string
	// 阿里体育openId
	_openId string
	// 合作方openId
	_thirdOpenId string
}

// NewAlibabaAlisportsPassportAuthUnbindRequest 初始化AlibabaAlisportsPassportAuthUnbindAPIRequest对象
func NewAlibabaAlisportsPassportAuthUnbindRequest() *AlibabaAlisportsPassportAuthUnbindAPIRequest {
	return &AlibabaAlisportsPassportAuthUnbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportAuthUnbindAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.passport.auth.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportAuthUnbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlisportsPassportAuthUnbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlispAppKey is AlispAppKey Setter
// 阿里体育业务KEY
func (r *AlibabaAlisportsPassportAuthUnbindAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaAlisportsPassportAuthUnbindAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetOpenId is OpenId Setter
// 阿里体育openId
func (r *AlibabaAlisportsPassportAuthUnbindAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r AlibabaAlisportsPassportAuthUnbindAPIRequest) GetOpenId() string {
	return r._openId
}

// SetThirdOpenId is ThirdOpenId Setter
// 合作方openId
func (r *AlibabaAlisportsPassportAuthUnbindAPIRequest) SetThirdOpenId(_thirdOpenId string) error {
	r._thirdOpenId = _thirdOpenId
	r.Set("third_open_id", _thirdOpenId)
	return nil
}

// GetThirdOpenId ThirdOpenId Getter
func (r AlibabaAlisportsPassportAuthUnbindAPIRequest) GetThirdOpenId() string {
	return r._thirdOpenId
}
