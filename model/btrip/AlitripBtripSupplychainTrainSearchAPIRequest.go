package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainTrainSearchAPIRequest 【商旅】火车票订单查询 API请求
// alitrip.btrip.supplychain.train.search
//
// 【商旅】火车票订单查询
type AlitripBtripSupplychainTrainSearchAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenApiSearchRq
}

// NewAlitripBtripSupplychainTrainSearchRequest 初始化AlitripBtripSupplychainTrainSearchAPIRequest对象
func NewAlitripBtripSupplychainTrainSearchRequest() *AlitripBtripSupplychainTrainSearchAPIRequest {
	return &AlitripBtripSupplychainTrainSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainTrainSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.train.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainTrainSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripSupplychainTrainSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripSupplychainTrainSearchAPIRequest) SetRq(_rq *OpenApiSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripSupplychainTrainSearchAPIRequest) GetRq() *OpenApiSearchRq {
	return r._rq
}
