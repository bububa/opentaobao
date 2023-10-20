package flightuppc

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
