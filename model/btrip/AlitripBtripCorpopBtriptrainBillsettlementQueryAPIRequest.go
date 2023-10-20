package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcorpopbtriptrainbillsettlementqueryAPIRequest 商旅火车票结算记账查询接口 API请求
// alitrip.btrip.corpop.btriptrain.billsettlement.query
//
// 商旅火车票结算记账查询接口
type AlitripbtripcorpopbtriptrainbillsettlementqueryAPIRequest struct {
	model.Params
	// 请求入参
	_rq *OpenIsvBillSettlementSearchRq
}

// NewAlitripbtripcorpopbtriptrainbillsettlementqueryRequest 初始化AlitripbtripcorpopbtriptrainbillsettlementqueryAPIRequest对象
func NewAlitripbtripcorpopbtriptrainbillsettlementqueryRequest() *AlitripbtripcorpopbtriptrainbillsettlementqueryAPIRequest {
	return &AlitripbtripcorpopbtriptrainbillsettlementqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcorpopbtriptrainbillsettlementqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.corpop.btriptrain.billsettlement.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcorpopbtriptrainbillsettlementqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcorpopbtriptrainbillsettlementqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求入参
func (r *AlitripbtripcorpopbtriptrainbillsettlementqueryAPIRequest) SetRq(_rq *OpenIsvBillSettlementSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcorpopbtriptrainbillsettlementqueryAPIRequest) GetRq() *OpenIsvBillSettlementSearchRq {
	return r._rq
}
