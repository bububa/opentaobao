package topoaid

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopSecretExtendAPIRequest 虚拟号延期 API请求
// taobao.top.secret.extend
//
// 虚拟号延期
type TaobaoTopSecretExtendAPIRequest struct {
	model.Params
	// 虚拟号延期请求
	_extendRequest *SecretNoExtendRequest
}

// NewTaobaoTopSecretExtendRequest 初始化TaobaoTopSecretExtendAPIRequest对象
func NewTaobaoTopSecretExtendRequest() *TaobaoTopSecretExtendAPIRequest {
	return &TaobaoTopSecretExtendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopSecretExtendAPIRequest) GetApiMethodName() string {
	return "taobao.top.secret.extend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopSecretExtendAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetExtendRequest is ExtendRequest Setter
// 虚拟号延期请求
func (r *TaobaoTopSecretExtendAPIRequest) SetExtendRequest(_extendRequest *SecretNoExtendRequest) error {
	r._extendRequest = _extendRequest
	r.Set("extend_request", _extendRequest)
	return nil
}

// GetExtendRequest ExtendRequest Getter
func (r TaobaoTopSecretExtendAPIRequest) GetExtendRequest() *SecretNoExtendRequest {
	return r._extendRequest
}
