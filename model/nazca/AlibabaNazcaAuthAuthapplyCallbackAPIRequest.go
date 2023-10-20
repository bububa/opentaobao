package nazca

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabanazcaauthauthapplycallbackAPIRequest 认证的统一回调接口 API请求
// alibaba.nazca.auth.authapply.callback
//
// 认证的统一回调接口
type AlibabanazcaauthauthapplycallbackAPIRequest struct {
	model.Params
	// 认证回调参数
	_authApplyDoneCallbackDo *AuthApplyDoneCallBackDo
}

// NewAlibabanazcaauthauthapplycallbackRequest 初始化AlibabanazcaauthauthapplycallbackAPIRequest对象
func NewAlibabanazcaauthauthapplycallbackRequest() *AlibabanazcaauthauthapplycallbackAPIRequest {
	return &AlibabanazcaauthauthapplycallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabanazcaauthauthapplycallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.nazca.auth.authapply.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabanazcaauthauthapplycallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabanazcaauthauthapplycallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuthApplyDoneCallbackDo is AuthApplyDoneCallbackDo Setter
// 认证回调参数
func (r *AlibabanazcaauthauthapplycallbackAPIRequest) SetAuthApplyDoneCallbackDo(_authApplyDoneCallbackDo *AuthApplyDoneCallBackDo) error {
	r._authApplyDoneCallbackDo = _authApplyDoneCallbackDo
	r.Set("auth_apply_done_callback_do", _authApplyDoneCallbackDo)
	return nil
}

// GetAuthApplyDoneCallbackDo AuthApplyDoneCallbackDo Getter
func (r AlibabanazcaauthauthapplycallbackAPIRequest) GetAuthApplyDoneCallbackDo() *AuthApplyDoneCallBackDo {
	return r._authApplyDoneCallbackDo
}
