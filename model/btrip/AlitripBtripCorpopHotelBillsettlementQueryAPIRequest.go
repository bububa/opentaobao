package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpophotelbillsettlementqueryAPIRequest 酒店结算记账查询接口 API请求
// alitrip.btrip.corpop.hotel.billsettlement.query
//
// 酒店结算记账查询接口
type AlitripbtripcorpophotelbillsettlementqueryAPIRequest struct {
	model.Params
	// 请求入参
	_rq *OpenIsvBillSettlementSearchRq
}

// NewAlitripbtripcorpophotelbillsettlementqueryRequest 初始化AlitripbtripcorpophotelbillsettlementqueryAPIRequest对象
func NewAlitripbtripcorpophotelbillsettlementqueryRequest() *AlitripbtripcorpophotelbillsettlementqueryAPIRequest {
	return &AlitripbtripcorpophotelbillsettlementqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcorpophotelbillsettlementqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.hotel.billsettlement.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcorpophotelbillsettlementqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcorpophotelbillsettlementqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求入参
func (r *AlitripbtripcorpophotelbillsettlementqueryAPIRequest) SetRq(_rq *OpenIsvBillSettlementSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcorpophotelbillsettlementqueryAPIRequest) GetRq() *OpenIsvBillSettlementSearchRq {
	return r._rq
}
