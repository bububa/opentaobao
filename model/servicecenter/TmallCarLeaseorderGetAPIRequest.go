package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseorderGetAPIRequest 获取租赁订单信息 API请求
// tmall.car.leaseorder.get
//
// 卖家通过供销平台获取代销商的订单信息，但是部分情况下网商银行订单号获取不到，需要提供接口或者工具给卖家
type TmallCarLeaseorderGetAPIRequest struct {
	model.Params
	// 订单号
	_orderId int64
}

// NewTmallCarLeaseorderGetRequest 初始化TmallCarLeaseorderGetAPIRequest对象
func NewTmallCarLeaseorderGetRequest() *TmallCarLeaseorderGetAPIRequest {
	return &TmallCarLeaseorderGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCarLeaseorderGetAPIRequest) Reset() {
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeaseorderGetAPIRequest) GetApiMethodName() string {
	return "tmall.car.leaseorder.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeaseorderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarLeaseorderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单号
func (r *TmallCarLeaseorderGetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TmallCarLeaseorderGetAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolTmallCarLeaseorderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCarLeaseorderGetRequest()
	},
}

// GetTmallCarLeaseorderGetRequest 从 sync.Pool 获取 TmallCarLeaseorderGetAPIRequest
func GetTmallCarLeaseorderGetAPIRequest() *TmallCarLeaseorderGetAPIRequest {
	return poolTmallCarLeaseorderGetAPIRequest.Get().(*TmallCarLeaseorderGetAPIRequest)
}

// ReleaseTmallCarLeaseorderGetAPIRequest 将 TmallCarLeaseorderGetAPIRequest 放入 sync.Pool
func ReleaseTmallCarLeaseorderGetAPIRequest(v *TmallCarLeaseorderGetAPIRequest) {
	v.Reset()
	poolTmallCarLeaseorderGetAPIRequest.Put(v)
}
