package mos

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosCommonAuthOperatorInfoAPIRequest 获取当前人员信息 API请求
// alibaba.mos.common.auth.operator.info
//
// 获取当前人员信息
type AlibabaMosCommonAuthOperatorInfoAPIRequest struct {
	model.Params
	// 登录授权后返回的token信息
	_token string
}

// NewAlibabaMosCommonAuthOperatorInfoRequest 初始化AlibabaMosCommonAuthOperatorInfoAPIRequest对象
func NewAlibabaMosCommonAuthOperatorInfoRequest() *AlibabaMosCommonAuthOperatorInfoAPIRequest {
	return &AlibabaMosCommonAuthOperatorInfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMosCommonAuthOperatorInfoAPIRequest) Reset() {
	r._token = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosCommonAuthOperatorInfoAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.common.auth.operator.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosCommonAuthOperatorInfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosCommonAuthOperatorInfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 登录授权后返回的token信息
func (r *AlibabaMosCommonAuthOperatorInfoAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaMosCommonAuthOperatorInfoAPIRequest) GetToken() string {
	return r._token
}

var poolAlibabaMosCommonAuthOperatorInfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMosCommonAuthOperatorInfoRequest()
	},
}

// GetAlibabaMosCommonAuthOperatorInfoRequest 从 sync.Pool 获取 AlibabaMosCommonAuthOperatorInfoAPIRequest
func GetAlibabaMosCommonAuthOperatorInfoAPIRequest() *AlibabaMosCommonAuthOperatorInfoAPIRequest {
	return poolAlibabaMosCommonAuthOperatorInfoAPIRequest.Get().(*AlibabaMosCommonAuthOperatorInfoAPIRequest)
}

// ReleaseAlibabaMosCommonAuthOperatorInfoAPIRequest 将 AlibabaMosCommonAuthOperatorInfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaMosCommonAuthOperatorInfoAPIRequest(v *AlibabaMosCommonAuthOperatorInfoAPIRequest) {
	v.Reset()
	poolAlibabaMosCommonAuthOperatorInfoAPIRequest.Put(v)
}
