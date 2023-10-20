package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabanazcaauthchangeauthapplycallbackAPIRequest 变更认证回调 API请求
// alibaba.nazca.auth.changeauthapply.callback
//
// 变更认证回调
type AlibabanazcaauthchangeauthapplycallbackAPIRequest struct {
	model.Params
	// 变更认证回调参数
	_paramChangeAuthCallBackDo *ChangeAuthCallBackDo
}

// NewAlibabanazcaauthchangeauthapplycallbackRequest 初始化AlibabanazcaauthchangeauthapplycallbackAPIRequest对象
func NewAlibabanazcaauthchangeauthapplycallbackRequest() *AlibabanazcaauthchangeauthapplycallbackAPIRequest {
	return &AlibabanazcaauthchangeauthapplycallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabanazcaauthchangeauthapplycallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.nazca.auth.changeauthapply.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabanazcaauthchangeauthapplycallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabanazcaauthchangeauthapplycallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamChangeAuthCallBackDo is ParamChangeAuthCallBackDo Setter
// 变更认证回调参数
func (r *AlibabanazcaauthchangeauthapplycallbackAPIRequest) SetParamChangeAuthCallBackDo(_paramChangeAuthCallBackDo *ChangeAuthCallBackDo) error {
	r._paramChangeAuthCallBackDo = _paramChangeAuthCallBackDo
	r.Set("param_change_auth_call_back_do", _paramChangeAuthCallBackDo)
	return nil
}

// GetParamChangeAuthCallBackDo ParamChangeAuthCallBackDo Getter
func (r AlibabanazcaauthchangeauthapplycallbackAPIRequest) GetParamChangeAuthCallBackDo() *ChangeAuthCallBackDo {
	return r._paramChangeAuthCallBackDo
}
