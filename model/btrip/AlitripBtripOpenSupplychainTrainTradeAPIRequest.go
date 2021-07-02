package btrip

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenSupplychainTrainTradeAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.supplychain.train.trade"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenSupplychainTrainTradeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
