package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymSteamShadowuserGetAPIRequest 获取影子标识 API请求
// alibaba.jym.steam.shadowuser.get
//
// 交易猫Steam类目获取影子ID
type AlibabaJymSteamShadowuserGetAPIRequest struct {
	model.Params
	// shadowToken
	_token string
}

// NewAlibabaJymSteamShadowuserGetRequest 初始化AlibabaJymSteamShadowuserGetAPIRequest对象
func NewAlibabaJymSteamShadowuserGetRequest() *AlibabaJymSteamShadowuserGetAPIRequest {
	return &AlibabaJymSteamShadowuserGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymSteamShadowuserGetAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.steam.shadowuser.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymSteamShadowuserGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymSteamShadowuserGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// shadowToken
func (r *AlibabaJymSteamShadowuserGetAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaJymSteamShadowuserGetAPIRequest) GetToken() string {
	return r._token
}
