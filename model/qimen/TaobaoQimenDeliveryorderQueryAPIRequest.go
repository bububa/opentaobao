package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenDeliveryorderQueryAPIRequest 发货单查询接口 API请求
// taobao.qimen.deliveryorder.query
//
// ERP调用奇门的发货单查询接口，查询发货单详情
type TaobaoQimenDeliveryorderQueryAPIRequest struct {
	model.Params
	//
	_request *DeliveryOrderQueryRequest
}

// NewTaobaoQimenDeliveryorderQueryRequest 初始化TaobaoQimenDeliveryorderQueryAPIRequest对象
func NewTaobaoQimenDeliveryorderQueryRequest() *TaobaoQimenDeliveryorderQueryAPIRequest {
	return &TaobaoQimenDeliveryorderQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenDeliveryorderQueryAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenDeliveryorderQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.deliveryorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenDeliveryorderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenDeliveryorderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenDeliveryorderQueryAPIRequest) SetRequest(_request *DeliveryOrderQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenDeliveryorderQueryAPIRequest) GetRequest() *DeliveryOrderQueryRequest {
	return r._request
}

var poolTaobaoQimenDeliveryorderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenDeliveryorderQueryRequest()
	},
}

// GetTaobaoQimenDeliveryorderQueryRequest 从 sync.Pool 获取 TaobaoQimenDeliveryorderQueryAPIRequest
func GetTaobaoQimenDeliveryorderQueryAPIRequest() *TaobaoQimenDeliveryorderQueryAPIRequest {
	return poolTaobaoQimenDeliveryorderQueryAPIRequest.Get().(*TaobaoQimenDeliveryorderQueryAPIRequest)
}

// ReleaseTaobaoQimenDeliveryorderQueryAPIRequest 将 TaobaoQimenDeliveryorderQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenDeliveryorderQueryAPIRequest(v *TaobaoQimenDeliveryorderQueryAPIRequest) {
	v.Reset()
	poolTaobaoQimenDeliveryorderQueryAPIRequest.Put(v)
}
