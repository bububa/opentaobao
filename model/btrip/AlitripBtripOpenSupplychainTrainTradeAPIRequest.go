package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripopensupplychaintraintradeAPIRequest 商旅火车票交易流水接口 API请求
// alitrip.btrip.open.supplychain.train.trade
//
// 商旅火车票交易流水接口
type AlitripbtripopensupplychaintraintradeAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenApiZzdSearchRq
}

// NewAlitripbtripopensupplychaintraintradeRequest 初始化AlitripbtripopensupplychaintraintradeAPIRequest对象
func NewAlitripbtripopensupplychaintraintradeRequest() *AlitripbtripopensupplychaintraintradeAPIRequest {
	return &AlitripbtripopensupplychaintraintradeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripopensupplychaintraintradeAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.supplychain.train.trade"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripopensupplychaintraintradeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripopensupplychaintraintradeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripbtripopensupplychaintraintradeAPIRequest) SetRq(_rq *OpenApiZzdSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripopensupplychaintraintradeAPIRequest) GetRq() *OpenApiZzdSearchRq {
	return r._rq
}
