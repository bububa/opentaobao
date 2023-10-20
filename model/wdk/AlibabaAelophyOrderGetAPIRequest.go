package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAelophyOrderGetAPIRequest 翱象拉取订单接口 API请求
// alibaba.aelophy.order.get
//
// 翱象拉取订单接口
type AlibabaAelophyOrderGetAPIRequest struct {
	model.Params
	// 请求对象
	_orderGetRequest *OrderGetRequest
}

// NewAlibabaAelophyOrderGetRequest 初始化AlibabaAelophyOrderGetAPIRequest对象
func NewAlibabaAelophyOrderGetRequest() *AlibabaAelophyOrderGetAPIRequest {
	return &AlibabaAelophyOrderGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAelophyOrderGetAPIRequest) Reset() {
	r._orderGetRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAelophyOrderGetAPIRequest) GetApiMethodName() string {
	return "alibaba.aelophy.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAelophyOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAelophyOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderGetRequest is OrderGetRequest Setter
// 请求对象
func (r *AlibabaAelophyOrderGetAPIRequest) SetOrderGetRequest(_orderGetRequest *OrderGetRequest) error {
	r._orderGetRequest = _orderGetRequest
	r.Set("order_get_request", _orderGetRequest)
	return nil
}

// GetOrderGetRequest OrderGetRequest Getter
func (r AlibabaAelophyOrderGetAPIRequest) GetOrderGetRequest() *OrderGetRequest {
	return r._orderGetRequest
}

var poolAlibabaAelophyOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAelophyOrderGetRequest()
	},
}

// GetAlibabaAelophyOrderGetRequest 从 sync.Pool 获取 AlibabaAelophyOrderGetAPIRequest
func GetAlibabaAelophyOrderGetAPIRequest() *AlibabaAelophyOrderGetAPIRequest {
	return poolAlibabaAelophyOrderGetAPIRequest.Get().(*AlibabaAelophyOrderGetAPIRequest)
}

// ReleaseAlibabaAelophyOrderGetAPIRequest 将 AlibabaAelophyOrderGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAelophyOrderGetAPIRequest(v *AlibabaAelophyOrderGetAPIRequest) {
	v.Reset()
	poolAlibabaAelophyOrderGetAPIRequest.Put(v)
}
