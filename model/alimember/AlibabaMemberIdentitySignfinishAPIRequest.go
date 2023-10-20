package alimember

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamemberidentitysignfinishAPIRequest 签约确认 API请求
// alibaba.member.identity.signfinish
//
// 签约确认
type AlibabamemberidentitysignfinishAPIRequest struct {
	model.Params
	// 签约确认信息
	_signFinish *SignIdentityFinishRequest
}

// NewAlibabamemberidentitysignfinishRequest 初始化AlibabamemberidentitysignfinishAPIRequest对象
func NewAlibabamemberidentitysignfinishRequest() *AlibabamemberidentitysignfinishAPIRequest {
	return &AlibabamemberidentitysignfinishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamemberidentitysignfinishAPIRequest) GetApiMethodName() string {
	return "alibaba.member.identity.signfinish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamemberidentitysignfinishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamemberidentitysignfinishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSignFinish is SignFinish Setter
// 签约确认信息
func (r *AlibabamemberidentitysignfinishAPIRequest) SetSignFinish(_signFinish *SignIdentityFinishRequest) error {
	r._signFinish = _signFinish
	r.Set("sign_finish", _signFinish)
	return nil
}

// GetSignFinish SignFinish Getter
func (r AlibabamemberidentitysignfinishAPIRequest) GetSignFinish() *SignIdentityFinishRequest {
	return r._signFinish
}
