package btrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripCorpopCarBillsettlementQueryAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopCarBillsettlementQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.car.billsettlement.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopCarBillsettlementQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCorpopCarBillsettlementQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripBtripCorpopCarBillsettlementQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripCorpopCarBillsettlementQueryRequest()
	},
}

// GetAlitripBtripCorpopCarBillsettlementQueryRequest 从 sync.Pool 获取 AlitripBtripCorpopCarBillsettlementQueryAPIRequest
func GetAlitripBtripCorpopCarBillsettlementQueryAPIRequest() *AlitripBtripCorpopCarBillsettlementQueryAPIRequest {
	return poolAlitripBtripCorpopCarBillsettlementQueryAPIRequest.Get().(*AlitripBtripCorpopCarBillsettlementQueryAPIRequest)
}

// ReleaseAlitripBtripCorpopCarBillsettlementQueryAPIRequest 将 AlitripBtripCorpopCarBillsettlementQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripCorpopCarBillsettlementQueryAPIRequest(v *AlitripBtripCorpopCarBillsettlementQueryAPIRequest) {
	v.Reset()
	poolAlitripBtripCorpopCarBillsettlementQueryAPIRequest.Put(v)
}
