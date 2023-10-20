package jym

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymsteamshadowusergetAPIRequest 获取影子标识 API请求
// alibaba.jym.steam.shadowuser.get
//
// 交易猫Steam类目获取影子ID
type AlibabajymsteamshadowusergetAPIRequest struct {
	model.Params
	// shadowToken
	_token string
}

// NewAlibabajymsteamshadowusergetRequest 初始化AlibabajymsteamshadowusergetAPIRequest对象
func NewAlibabajymsteamshadowusergetRequest() *AlibabajymsteamshadowusergetAPIRequest {
	return &AlibabajymsteamshadowusergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymsteamshadowusergetAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.steam.shadowuser.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymsteamshadowusergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymsteamshadowusergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// shadowToken
func (r *AlibabajymsteamshadowusergetAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabajymsteamshadowusergetAPIRequest) GetToken() string {
	return r._token
}
