package nazca

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest 变更认证回调 API请求
// alibaba.nazca.auth.changeauthapply.callback
//
// 变更认证回调
type AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest struct {
	model.Params
	// 变更认证回调参数
	_paramChangeAuthCallBackDo *ChangeAuthCallBackDo
}

// NewAlibabaNazcaAuthChangeauthapplyCallbackRequest 初始化AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest对象
func NewAlibabaNazcaAuthChangeauthapplyCallbackRequest() *AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest {
	return &AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest) Reset() {
	r._paramChangeAuthCallBackDo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest) GetApiMethodName() string {
	return "alibaba.nazca.auth.changeauthapply.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamChangeAuthCallBackDo is ParamChangeAuthCallBackDo Setter
// 变更认证回调参数
func (r *AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest) SetParamChangeAuthCallBackDo(_paramChangeAuthCallBackDo *ChangeAuthCallBackDo) error {
	r._paramChangeAuthCallBackDo = _paramChangeAuthCallBackDo
	r.Set("param_change_auth_call_back_do", _paramChangeAuthCallBackDo)
	return nil
}

// GetParamChangeAuthCallBackDo ParamChangeAuthCallBackDo Getter
func (r AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest) GetParamChangeAuthCallBackDo() *ChangeAuthCallBackDo {
	return r._paramChangeAuthCallBackDo
}

var poolAlibabaNazcaAuthChangeauthapplyCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaNazcaAuthChangeauthapplyCallbackRequest()
	},
}

// GetAlibabaNazcaAuthChangeauthapplyCallbackRequest 从 sync.Pool 获取 AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest
func GetAlibabaNazcaAuthChangeauthapplyCallbackAPIRequest() *AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest {
	return poolAlibabaNazcaAuthChangeauthapplyCallbackAPIRequest.Get().(*AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest)
}

// ReleaseAlibabaNazcaAuthChangeauthapplyCallbackAPIRequest 将 AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest 放入 sync.Pool
func ReleaseAlibabaNazcaAuthChangeauthapplyCallbackAPIRequest(v *AlibabaNazcaAuthChangeauthapplyCallbackAPIRequest) {
	v.Reset()
	poolAlibabaNazcaAuthChangeauthapplyCallbackAPIRequest.Put(v)
}
