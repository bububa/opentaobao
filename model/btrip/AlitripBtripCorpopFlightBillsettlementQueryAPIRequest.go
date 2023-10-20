package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpopflightbillsettlementqueryAPIRequest 机票结算记账查询接口 API请求
// alitrip.btrip.corpop.flight.billsettlement.query
//
// 机票结算记账查询接口
type AlitripbtripcorpopflightbillsettlementqueryAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenIsvBillSettlementSearchRq
}

// NewAlitripbtripcorpopflightbillsettlementqueryRequest 初始化AlitripbtripcorpopflightbillsettlementqueryAPIRequest对象
func NewAlitripbtripcorpopflightbillsettlementqueryRequest() *AlitripbtripcorpopflightbillsettlementqueryAPIRequest {
	return &AlitripbtripcorpopflightbillsettlementqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcorpopflightbillsettlementqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.flight.billsettlement.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcorpopflightbillsettlementqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcorpopflightbillsettlementqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripbtripcorpopflightbillsettlementqueryAPIRequest) SetRq(_rq *OpenIsvBillSettlementSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcorpopflightbillsettlementqueryAPIRequest) GetRq() *OpenIsvBillSettlementSearchRq {
	return r._rq
}
