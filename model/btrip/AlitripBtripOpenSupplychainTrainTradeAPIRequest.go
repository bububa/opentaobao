package btrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenSupplychainTrainTradeAPIRequest 商旅火车票交易流水接口 API请求
// alitrip.btrip.open.supplychain.train.trade
//
// 商旅火车票交易流水接口
type AlitripBtripOpenSupplychainTrainTradeAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenApiZzdSearchRq
}

// NewAlitripBtripOpenSupplychainTrainTradeRequest 初始化AlitripBtripOpenSupplychainTrainTradeAPIRequest对象
func NewAlitripBtripOpenSupplychainTrainTradeRequest() *AlitripBtripOpenSupplychainTrainTradeAPIRequest {
	return &AlitripBtripOpenSupplychainTrainTradeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripBtripOpenSupplychainTrainTradeAPIRequest) Reset() {
	r._rq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenSupplychainTrainTradeAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.supplychain.train.trade"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenSupplychainTrainTradeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripOpenSupplychainTrainTradeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripBtripOpenSupplychainTrainTradeAPIRequest) SetRq(_rq *OpenApiZzdSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripOpenSupplychainTrainTradeAPIRequest) GetRq() *OpenApiZzdSearchRq {
	return r._rq
}

var poolAlitripBtripOpenSupplychainTrainTradeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripBtripOpenSupplychainTrainTradeRequest()
	},
}

// GetAlitripBtripOpenSupplychainTrainTradeRequest 从 sync.Pool 获取 AlitripBtripOpenSupplychainTrainTradeAPIRequest
func GetAlitripBtripOpenSupplychainTrainTradeAPIRequest() *AlitripBtripOpenSupplychainTrainTradeAPIRequest {
	return poolAlitripBtripOpenSupplychainTrainTradeAPIRequest.Get().(*AlitripBtripOpenSupplychainTrainTradeAPIRequest)
}

// ReleaseAlitripBtripOpenSupplychainTrainTradeAPIRequest 将 AlitripBtripOpenSupplychainTrainTradeAPIRequest 放入 sync.Pool
func ReleaseAlitripBtripOpenSupplychainTrainTradeAPIRequest(v *AlitripBtripOpenSupplychainTrainTradeAPIRequest) {
	v.Reset()
	poolAlitripBtripOpenSupplychainTrainTradeAPIRequest.Put(v)
}
