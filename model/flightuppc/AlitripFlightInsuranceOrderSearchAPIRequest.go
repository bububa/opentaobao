package flightuppc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripflightinsuranceordersearchAPIRequest 查询保险订单详情 API请求
// alitrip.flight.insurance.order.search
//
// 查询保险订单详情
type AlitripflightinsuranceordersearchAPIRequest struct {
	model.Params
	// 外部订单号
	_outOrderId int64
}

// NewAlitripflightinsuranceordersearchRequest 初始化AlitripflightinsuranceordersearchAPIRequest对象
func NewAlitripflightinsuranceordersearchRequest() *AlitripflightinsuranceordersearchAPIRequest {
	return &AlitripflightinsuranceordersearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripflightinsuranceordersearchAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.insurance.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripflightinsuranceordersearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripflightinsuranceordersearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutOrderId is OutOrderId Setter
// 外部订单号
func (r *AlitripflightinsuranceordersearchAPIRequest) SetOutOrderId(_outOrderId int64) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlitripflightinsuranceordersearchAPIRequest) GetOutOrderId() int64 {
	return r._outOrderId
}
