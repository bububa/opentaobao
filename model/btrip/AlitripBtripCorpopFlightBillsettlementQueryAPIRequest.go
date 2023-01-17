package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopFlightBillsettlementQueryAPIRequest 机票结算记账查询接口 API请求
// alitrip.btrip.corpop.flight.billsettlement.query
//
// 机票结算记账查询接口
type AlitripBtripCorpopFlightBillsettlementQueryAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenIsvBillSettlementSearchRq
}

// NewAlitripBtripCorpopFlightBillsettlementQueryRequest 初始化AlitripBtripCorpopFlightBillsettlementQueryAPIRequest对象
func NewAlitripBtripCorpopFlightBillsettlementQueryRequest() *AlitripBtripCorpopFlightBillsettlementQueryAPIRequest {
	return &AlitripBtripCorpopFlightBillsettlementQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopFlightBillsettlementQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.flight.billsettlement.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopFlightBillsettlementQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCorpopFlightBillsettlementQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripCorpopFlightBillsettlementQueryAPIRequest) SetRq(_rq *OpenIsvBillSettlementSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopFlightBillsettlementQueryAPIRequest) GetRq() *OpenIsvBillSettlementSearchRq {
	return r._rq
}
