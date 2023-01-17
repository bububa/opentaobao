package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest 【商旅】火车票订单详情查询 API请求
// alitrip.btrip.supplychain.train.detail.search.vtwo
//
// 【商旅】火车票订单详情查询
type AlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenApiSearchDetailRq
}

// NewAlitripBtripSupplychainTrainDetailSearchVtwoRequest 初始化AlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest对象
func NewAlitripBtripSupplychainTrainDetailSearchVtwoRequest() *AlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest {
	return &AlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.train.detail.search.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest) SetRq(_rq *OpenApiSearchDetailRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest) GetRq() *OpenApiSearchDetailRq {
	return r._rq
}
