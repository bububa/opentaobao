package btrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
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

var poolAlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripSupplychainTrainDetailSearchVtwoRequest()
	},
}

// GetAlitripBtripSupplychainTrainDetailSearchVtwoRequest 从 sync.Pool 获取 AlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest
func GetAlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest() *AlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest {
	return poolAlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest.Get().(*AlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest)
}

// ReleaseAlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest 将 AlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest(v *AlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest) {
	v.Reset()
	poolAlitripBtripSupplychainTrainDetailSearchVtwoAPIRequest.Put(v)
}
