package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripSupplychainTrainDetailSearchAPIRequest 【商旅】火车票订单详情查询 API请求
// alitrip.btrip.supplychain.train.detail.search
//
// 【商旅】火车票订单详情查询
type AlitripBtripSupplychainTrainDetailSearchAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenApiSearchDetailRq
}

// NewAlitripBtripSupplychainTrainDetailSearchRequest 初始化AlitripBtripSupplychainTrainDetailSearchAPIRequest对象
func NewAlitripBtripSupplychainTrainDetailSearchRequest() *AlitripBtripSupplychainTrainDetailSearchAPIRequest {
	return &AlitripBtripSupplychainTrainDetailSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainTrainDetailSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.train.detail.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainTrainDetailSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripBtripSupplychainTrainDetailSearchAPIRequest) SetRq(_rq *OpenApiSearchDetailRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripSupplychainTrainDetailSearchAPIRequest) GetRq() *OpenApiSearchDetailRq {
	return r._rq
}
