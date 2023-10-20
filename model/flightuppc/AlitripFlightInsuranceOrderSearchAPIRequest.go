package flightuppc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightInsuranceOrderSearchAPIRequest 查询保险订单详情 API请求
// alitrip.flight.insurance.order.search
//
// 查询保险订单详情
type AlitripFlightInsuranceOrderSearchAPIRequest struct {
	model.Params
	// 外部订单号
	_outOrderId int64
}

// NewAlitripFlightInsuranceOrderSearchRequest 初始化AlitripFlightInsuranceOrderSearchAPIRequest对象
func NewAlitripFlightInsuranceOrderSearchRequest() *AlitripFlightInsuranceOrderSearchAPIRequest {
	return &AlitripFlightInsuranceOrderSearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripFlightInsuranceOrderSearchAPIRequest) Reset() {
	r._outOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripFlightInsuranceOrderSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.insurance.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripFlightInsuranceOrderSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripFlightInsuranceOrderSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutOrderId is OutOrderId Setter
// 外部订单号
func (r *AlitripFlightInsuranceOrderSearchAPIRequest) SetOutOrderId(_outOrderId int64) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlitripFlightInsuranceOrderSearchAPIRequest) GetOutOrderId() int64 {
	return r._outOrderId
}

var poolAlitripFlightInsuranceOrderSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripFlightInsuranceOrderSearchRequest()
	},
}

// GetAlitripFlightInsuranceOrderSearchRequest 从 sync.Pool 获取 AlitripFlightInsuranceOrderSearchAPIRequest
func GetAlitripFlightInsuranceOrderSearchAPIRequest() *AlitripFlightInsuranceOrderSearchAPIRequest {
	return poolAlitripFlightInsuranceOrderSearchAPIRequest.Get().(*AlitripFlightInsuranceOrderSearchAPIRequest)
}

// ReleaseAlitripFlightInsuranceOrderSearchAPIRequest 将 AlitripFlightInsuranceOrderSearchAPIRequest 放入 sync.Pool
func ReleaseAlitripFlightInsuranceOrderSearchAPIRequest(v *AlitripFlightInsuranceOrderSearchAPIRequest) {
	v.Reset()
	poolAlitripFlightInsuranceOrderSearchAPIRequest.Put(v)
}
