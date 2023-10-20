package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamoscommonauthoperatorinfoAPIRequest 获取当前人员信息 API请求
// alibaba.mos.common.auth.operator.info
//
// 获取当前人员信息
type AlibabamoscommonauthoperatorinfoAPIRequest struct {
	model.Params
	// 登录授权后返回的token信息
	_token string
}

// NewAlibabamoscommonauthoperatorinfoRequest 初始化AlibabamoscommonauthoperatorinfoAPIRequest对象
func NewAlibabamoscommonauthoperatorinfoRequest() *AlibabamoscommonauthoperatorinfoAPIRequest {
	return &AlibabamoscommonauthoperatorinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamoscommonauthoperatorinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.common.auth.operator.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamoscommonauthoperatorinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamoscommonauthoperatorinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 登录授权后返回的token信息
func (r *AlibabamoscommonauthoperatorinfoAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabamoscommonauthoperatorinfoAPIRequest) GetToken() string {
	return r._token
}
