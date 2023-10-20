package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpopcarbillsettlementqueryAPIRequest 用车结算记账查询接口 API请求
// alitrip.btrip.corpop.car.billsettlement.query
//
// 用车结算记账查询接口
type AlitripbtripcorpopcarbillsettlementqueryAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenIsvBillSettlementSearchRq
}

// NewAlitripbtripcorpopcarbillsettlementqueryRequest 初始化AlitripbtripcorpopcarbillsettlementqueryAPIRequest对象
func NewAlitripbtripcorpopcarbillsettlementqueryRequest() *AlitripbtripcorpopcarbillsettlementqueryAPIRequest {
	return &AlitripbtripcorpopcarbillsettlementqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcorpopcarbillsettlementqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.car.billsettlement.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcorpopcarbillsettlementqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcorpopcarbillsettlementqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripbtripcorpopcarbillsettlementqueryAPIRequest) SetRq(_rq *OpenIsvBillSettlementSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcorpopcarbillsettlementqueryAPIRequest) GetRq() *OpenIsvBillSettlementSearchRq {
	return r._rq
}
