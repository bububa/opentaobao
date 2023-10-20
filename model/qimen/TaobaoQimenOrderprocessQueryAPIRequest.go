package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderprocessQueryAPIRequest 订单流水查询接口 API请求
// taobao.qimen.orderprocess.query
//
// ERP调用订单流水查询接口
type TaobaoQimenOrderprocessQueryAPIRequest struct {
	model.Params
	//
	_request *OrderProcessQueryRequest
}

// NewTaobaoQimenOrderprocessQueryRequest 初始化TaobaoQimenOrderprocessQueryAPIRequest对象
func NewTaobaoQimenOrderprocessQueryRequest() *TaobaoQimenOrderprocessQueryAPIRequest {
	return &TaobaoQimenOrderprocessQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenOrderprocessQueryAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderprocessQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.orderprocess.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderprocessQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenOrderprocessQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenOrderprocessQueryAPIRequest) SetRequest(_request *OrderProcessQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenOrderprocessQueryAPIRequest) GetRequest() *OrderProcessQueryRequest {
	return r._request
}

var poolTaobaoQimenOrderprocessQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenOrderprocessQueryRequest()
	},
}

// GetTaobaoQimenOrderprocessQueryRequest 从 sync.Pool 获取 TaobaoQimenOrderprocessQueryAPIRequest
func GetTaobaoQimenOrderprocessQueryAPIRequest() *TaobaoQimenOrderprocessQueryAPIRequest {
	return poolTaobaoQimenOrderprocessQueryAPIRequest.Get().(*TaobaoQimenOrderprocessQueryAPIRequest)
}

// ReleaseTaobaoQimenOrderprocessQueryAPIRequest 将 TaobaoQimenOrderprocessQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenOrderprocessQueryAPIRequest(v *TaobaoQimenOrderprocessQueryAPIRequest) {
	v.Reset()
	poolTaobaoQimenOrderprocessQueryAPIRequest.Put(v)
}
