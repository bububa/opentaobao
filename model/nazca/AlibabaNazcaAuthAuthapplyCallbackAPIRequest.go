package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNazcaAuthAuthapplyCallbackAPIRequest 认证的统一回调接口 API请求
// alibaba.nazca.auth.authapply.callback
//
// 认证的统一回调接口
type AlibabaNazcaAuthAuthapplyCallbackAPIRequest struct {
	model.Params
	// 认证回调参数
	_authApplyDoneCallbackDo *AuthApplyDoneCallBackDo
}

// NewAlibabaNazcaAuthAuthapplyCallbackRequest 初始化AlibabaNazcaAuthAuthapplyCallbackAPIRequest对象
func NewAlibabaNazcaAuthAuthapplyCallbackRequest() *AlibabaNazcaAuthAuthapplyCallbackAPIRequest {
	return &AlibabaNazcaAuthAuthapplyCallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNazcaAuthAuthapplyCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.nazca.auth.authapply.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNazcaAuthAuthapplyCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaNazcaAuthAuthapplyCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuthApplyDoneCallbackDo is AuthApplyDoneCallbackDo Setter
// 认证回调参数
func (r *AlibabaNazcaAuthAuthapplyCallbackAPIRequest) SetAuthApplyDoneCallbackDo(_authApplyDoneCallbackDo *AuthApplyDoneCallBackDo) error {
	r._authApplyDoneCallbackDo = _authApplyDoneCallbackDo
	r.Set("auth_apply_done_callback_do", _authApplyDoneCallbackDo)
	return nil
}

// GetAuthApplyDoneCallbackDo AuthApplyDoneCallbackDo Getter
func (r AlibabaNazcaAuthAuthapplyCallbackAPIRequest) GetAuthApplyDoneCallbackDo() *AuthApplyDoneCallBackDo {
	return r._authApplyDoneCallbackDo
}
