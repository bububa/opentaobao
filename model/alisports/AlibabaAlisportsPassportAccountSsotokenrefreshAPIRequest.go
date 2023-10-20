package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalisportspassportaccountssotokenrefreshAPIRequest sso_token刷新 API请求
// alibaba.alisports.passport.account.ssotokenrefresh
//
// sso_token刷新
type AlibabaalisportspassportaccountssotokenrefreshAPIRequest struct {
	model.Params
	// 应用appkey
	_alispAppKey string
	// 当前时间戳[精确到秒,10位]
	_alispTime string
	// 签名
	_alispSign string
	// 登录成功返回的access_token(access_token是TOP保留关键字，只能改名，望谅解)
	_secret string
}

// NewAlibabaalisportspassportaccountssotokenrefreshRequest 初始化AlibabaalisportspassportaccountssotokenrefreshAPIRequest对象
func NewAlibabaalisportspassportaccountssotokenrefreshRequest() *AlibabaalisportspassportaccountssotokenrefreshAPIRequest {
	return &AlibabaalisportspassportaccountssotokenrefreshAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalisportspassportaccountssotokenrefreshAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.passport.account.ssotokenrefresh"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalisportspassportaccountssotokenrefreshAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalisportspassportaccountssotokenrefreshAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlispAppKey is AlispAppKey Setter
// 应用appkey
func (r *AlibabaalisportspassportaccountssotokenrefreshAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaalisportspassportaccountssotokenrefreshAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetAlispTime is AlispTime Setter
// 当前时间戳[精确到秒,10位]
func (r *AlibabaalisportspassportaccountssotokenrefreshAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// GetAlispTime AlispTime Getter
func (r AlibabaalisportspassportaccountssotokenrefreshAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// SetAlispSign is AlispSign Setter
// 签名
func (r *AlibabaalisportspassportaccountssotokenrefreshAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// GetAlispSign AlispSign Getter
func (r AlibabaalisportspassportaccountssotokenrefreshAPIRequest) GetAlispSign() string {
	return r._alispSign
}

// SetSecret is Secret Setter
// 登录成功返回的access_token(access_token是TOP保留关键字，只能改名，望谅解)
func (r *AlibabaalisportspassportaccountssotokenrefreshAPIRequest) SetSecret(_secret string) error {
	r._secret = _secret
	r.Set("secret", _secret)
	return nil
}

// GetSecret Secret Getter
func (r AlibabaalisportspassportaccountssotokenrefreshAPIRequest) GetSecret() string {
	return r._secret
}
