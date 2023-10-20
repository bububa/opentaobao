package tbtrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopOnceTokenGetAPIRequest 网关一次性token获取 API请求
// taobao.top.once.token.get
//
// 网关一次性token获取，对接文档:
type TaobaoTopOnceTokenGetAPIRequest struct {
	model.Params
	// sec_token
	_secToken string
}

// NewTaobaoTopOnceTokenGetRequest 初始化TaobaoTopOnceTokenGetAPIRequest对象
func NewTaobaoTopOnceTokenGetRequest() *TaobaoTopOnceTokenGetAPIRequest {
	return &TaobaoTopOnceTokenGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTopOnceTokenGetAPIRequest) Reset() {
	r._secToken = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTopOnceTokenGetAPIRequest) GetApiMethodName() string {
	return "taobao.top.once.token.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTopOnceTokenGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTopOnceTokenGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSecToken is SecToken Setter
// sec_token
func (r *TaobaoTopOnceTokenGetAPIRequest) SetSecToken(_secToken string) error {
	r._secToken = _secToken
	r.Set("sec_token", _secToken)
	return nil
}

// GetSecToken SecToken Getter
func (r TaobaoTopOnceTokenGetAPIRequest) GetSecToken() string {
	return r._secToken
}

var poolTaobaoTopOnceTokenGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTopOnceTokenGetRequest()
	},
}

// GetTaobaoTopOnceTokenGetRequest 从 sync.Pool 获取 TaobaoTopOnceTokenGetAPIRequest
func GetTaobaoTopOnceTokenGetAPIRequest() *TaobaoTopOnceTokenGetAPIRequest {
	return poolTaobaoTopOnceTokenGetAPIRequest.Get().(*TaobaoTopOnceTokenGetAPIRequest)
}

// ReleaseTaobaoTopOnceTokenGetAPIRequest 将 TaobaoTopOnceTokenGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTopOnceTokenGetAPIRequest(v *TaobaoTopOnceTokenGetAPIRequest) {
	v.Reset()
	poolTaobaoTopOnceTokenGetAPIRequest.Put(v)
}
