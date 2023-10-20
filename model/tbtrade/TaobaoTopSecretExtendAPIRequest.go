package tbtrade

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTopSecretExtendAPIRequest) Reset() {
	r._extendRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopSecretExtendAPIRequest) GetApiMethodName() string {
	return "taobao.top.secret.extend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopSecretExtendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTopSecretExtendAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoTopSecretExtendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTopSecretExtendRequest()
	},
}

// GetTaobaoTopSecretExtendRequest 从 sync.Pool 获取 TaobaoTopSecretExtendAPIRequest
func GetTaobaoTopSecretExtendAPIRequest() *TaobaoTopSecretExtendAPIRequest {
	return poolTaobaoTopSecretExtendAPIRequest.Get().(*TaobaoTopSecretExtendAPIRequest)
}

// ReleaseTaobaoTopSecretExtendAPIRequest 将 TaobaoTopSecretExtendAPIRequest 放入 sync.Pool
func ReleaseTaobaoTopSecretExtendAPIRequest(v *TaobaoTopSecretExtendAPIRequest) {
	v.Reset()
	poolTaobaoTopSecretExtendAPIRequest.Put(v)
}
