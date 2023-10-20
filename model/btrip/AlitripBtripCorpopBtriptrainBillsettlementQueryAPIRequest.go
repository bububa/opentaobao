package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest 商旅火车票结算记账查询接口 API请求
// alitrip.btrip.corpop.btriptrain.billsettlement.query
//
// 商旅火车票结算记账查询接口
type AlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest struct {
	model.Params
	// 请求入参
	_rq *OpenIsvBillSettlementSearchRq
}

// NewAlitripBtripCorpopBtriptrainBillsettlementQueryRequest 初始化AlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest对象
func NewAlitripBtripCorpopBtriptrainBillsettlementQueryRequest() *AlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest {
	return &AlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.btriptrain.billsettlement.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求入参
func (r *AlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest) SetRq(_rq *OpenIsvBillSettlementSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest) GetRq() *OpenIsvBillSettlementSearchRq {
	return r._rq
}

var poolAlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripCorpopBtriptrainBillsettlementQueryRequest()
	},
}

// GetAlitripBtripCorpopBtriptrainBillsettlementQueryRequest 从 sync.Pool 获取 AlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest
func GetAlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest() *AlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest {
	return poolAlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest.Get().(*AlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest)
}

// ReleaseAlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest 将 AlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest(v *AlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest) {
	v.Reset()
	poolAlitripBtripCorpopBtriptrainBillsettlementQueryAPIRequest.Put(v)
}
