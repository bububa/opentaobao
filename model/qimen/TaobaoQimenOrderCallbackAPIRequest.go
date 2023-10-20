package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderCallbackAPIRequest 配送拦截接口 API请求
// taobao.qimen.order.callback
//
// 配送拦截
type TaobaoQimenOrderCallbackAPIRequest struct {
	model.Params
	//
	_request *OrderCallbackRequestDo
}

// NewTaobaoQimenOrderCallbackRequest 初始化TaobaoQimenOrderCallbackAPIRequest对象
func NewTaobaoQimenOrderCallbackRequest() *TaobaoQimenOrderCallbackAPIRequest {
	return &TaobaoQimenOrderCallbackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenOrderCallbackAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderCallbackAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.order.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenOrderCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenOrderCallbackAPIRequest) SetRequest(_request *OrderCallbackRequestDo) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenOrderCallbackAPIRequest) GetRequest() *OrderCallbackRequestDo {
	return r._request
}

var poolTaobaoQimenOrderCallbackAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenOrderCallbackRequest()
	},
}

// GetTaobaoQimenOrderCallbackRequest 从 sync.Pool 获取 TaobaoQimenOrderCallbackAPIRequest
func GetTaobaoQimenOrderCallbackAPIRequest() *TaobaoQimenOrderCallbackAPIRequest {
	return poolTaobaoQimenOrderCallbackAPIRequest.Get().(*TaobaoQimenOrderCallbackAPIRequest)
}

// ReleaseTaobaoQimenOrderCallbackAPIRequest 将 TaobaoQimenOrderCallbackAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenOrderCallbackAPIRequest(v *TaobaoQimenOrderCallbackAPIRequest) {
	v.Reset()
	poolTaobaoQimenOrderCallbackAPIRequest.Put(v)
}
