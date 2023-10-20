package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderCancelAPIRequest 单据取消接口 API请求
// taobao.qimen.order.cancel
//
// taobao.qimen.order.cancel
type TaobaoQimenOrderCancelAPIRequest struct {
	model.Params
	//
	_request *OrderCancelRequest
}

// NewTaobaoQimenOrderCancelRequest 初始化TaobaoQimenOrderCancelAPIRequest对象
func NewTaobaoQimenOrderCancelRequest() *TaobaoQimenOrderCancelAPIRequest {
	return &TaobaoQimenOrderCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenOrderCancelAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderCancelAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenOrderCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenOrderCancelAPIRequest) SetRequest(_request *OrderCancelRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenOrderCancelAPIRequest) GetRequest() *OrderCancelRequest {
	return r._request
}

var poolTaobaoQimenOrderCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenOrderCancelRequest()
	},
}

// GetTaobaoQimenOrderCancelRequest 从 sync.Pool 获取 TaobaoQimenOrderCancelAPIRequest
func GetTaobaoQimenOrderCancelAPIRequest() *TaobaoQimenOrderCancelAPIRequest {
	return poolTaobaoQimenOrderCancelAPIRequest.Get().(*TaobaoQimenOrderCancelAPIRequest)
}

// ReleaseTaobaoQimenOrderCancelAPIRequest 将 TaobaoQimenOrderCancelAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenOrderCancelAPIRequest(v *TaobaoQimenOrderCancelAPIRequest) {
	v.Reset()
	poolTaobaoQimenOrderCancelAPIRequest.Put(v)
}
