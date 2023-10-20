package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoRefundRefundactionsGetAPIRequest 判断该订单能执行的逆向操作集合列表 API请求
// cainiao.refund.refundactions.get
//
// 判断该订单能执行的逆向操作集合列表
type CainiaoRefundRefundactionsGetAPIRequest struct {
	model.Params
	// 子订单ID
	_orderId string
}

// NewCainiaoRefundRefundactionsGetRequest 初始化CainiaoRefundRefundactionsGetAPIRequest对象
func NewCainiaoRefundRefundactionsGetRequest() *CainiaoRefundRefundactionsGetAPIRequest {
	return &CainiaoRefundRefundactionsGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoRefundRefundactionsGetAPIRequest) Reset() {
	r._orderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoRefundRefundactionsGetAPIRequest) GetApiMethodName() string {
	return "cainiao.refund.refundactions.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoRefundRefundactionsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoRefundRefundactionsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 子订单ID
func (r *CainiaoRefundRefundactionsGetAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r CainiaoRefundRefundactionsGetAPIRequest) GetOrderId() string {
	return r._orderId
}

var poolCainiaoRefundRefundactionsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoRefundRefundactionsGetRequest()
	},
}

// GetCainiaoRefundRefundactionsGetRequest 从 sync.Pool 获取 CainiaoRefundRefundactionsGetAPIRequest
func GetCainiaoRefundRefundactionsGetAPIRequest() *CainiaoRefundRefundactionsGetAPIRequest {
	return poolCainiaoRefundRefundactionsGetAPIRequest.Get().(*CainiaoRefundRefundactionsGetAPIRequest)
}

// ReleaseCainiaoRefundRefundactionsGetAPIRequest 将 CainiaoRefundRefundactionsGetAPIRequest 放入 sync.Pool
func ReleaseCainiaoRefundRefundactionsGetAPIRequest(v *CainiaoRefundRefundactionsGetAPIRequest) {
	v.Reset()
	poolCainiaoRefundRefundactionsGetAPIRequest.Put(v)
}
