package tbtrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotopsecretextendAPIRequest 虚拟号延期 API请求
// taobao.top.secret.extend
//
// 虚拟号延期
type TaobaotopsecretextendAPIRequest struct {
	model.Params
	// 虚拟号延期请求
	_extendRequest *SecretNoExtendRequest
}

// NewTaobaotopsecretextendRequest 初始化TaobaotopsecretextendAPIRequest对象
func NewTaobaotopsecretextendRequest() *TaobaotopsecretextendAPIRequest {
	return &TaobaotopsecretextendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotopsecretextendAPIRequest) GetApiMethodName() string {
	return "taobao.top.secret.extend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotopsecretextendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotopsecretextendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtendRequest is ExtendRequest Setter
// 虚拟号延期请求
func (r *TaobaotopsecretextendAPIRequest) SetExtendRequest(_extendRequest *SecretNoExtendRequest) error {
	r._extendRequest = _extendRequest
	r.Set("extend_request", _extendRequest)
	return nil
}

// GetExtendRequest ExtendRequest Getter
func (r TaobaotopsecretextendAPIRequest) GetExtendRequest() *SecretNoExtendRequest {
	return r._extendRequest
}
