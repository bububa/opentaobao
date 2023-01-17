package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopHotelBillsettlementQueryAPIRequest 酒店结算记账查询接口 API请求
// alitrip.btrip.corpop.hotel.billsettlement.query
//
// 酒店结算记账查询接口
type AlitripBtripCorpopHotelBillsettlementQueryAPIRequest struct {
	model.Params
	// 请求入参
	_rq *OpenIsvBillSettlementSearchRq
}

// NewAlitripBtripCorpopHotelBillsettlementQueryRequest 初始化AlitripBtripCorpopHotelBillsettlementQueryAPIRequest对象
func NewAlitripBtripCorpopHotelBillsettlementQueryRequest() *AlitripBtripCorpopHotelBillsettlementQueryAPIRequest {
	return &AlitripBtripCorpopHotelBillsettlementQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopHotelBillsettlementQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.hotel.billsettlement.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopHotelBillsettlementQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCorpopHotelBillsettlementQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求入参
func (r *AlitripBtripCorpopHotelBillsettlementQueryAPIRequest) SetRq(_rq *OpenIsvBillSettlementSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopHotelBillsettlementQueryAPIRequest) GetRq() *OpenIsvBillSettlementSearchRq {
	return r._rq
}
