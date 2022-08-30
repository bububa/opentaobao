package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopCarBillsettlementQueryAPIRequest 用车结算记账查询接口 API请求
// alitrip.btrip.corpop.car.billsettlement.query
//
// 用车结算记账查询接口
type AlitripBtripCorpopCarBillsettlementQueryAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenIsvBillSettlementSearchRq
}

// NewAlitripBtripCorpopCarBillsettlementQueryRequest 初始化AlitripBtripCorpopCarBillsettlementQueryAPIRequest对象
func NewAlitripBtripCorpopCarBillsettlementQueryRequest() *AlitripBtripCorpopCarBillsettlementQueryAPIRequest {
	return &AlitripBtripCorpopCarBillsettlementQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopCarBillsettlementQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.car.billsettlement.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopCarBillsettlementQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripCorpopCarBillsettlementQueryAPIRequest) SetRq(_rq *OpenIsvBillSettlementSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopCarBillsettlementQueryAPIRequest) GetRq() *OpenIsvBillSettlementSearchRq {
	return r._rq
}
