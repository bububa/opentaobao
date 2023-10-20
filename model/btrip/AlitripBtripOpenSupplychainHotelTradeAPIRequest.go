package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripopensupplychainhoteltradeAPIRequest 【商旅】酒店交易查询流水接口 API请求
// alitrip.btrip.open.supplychain.hotel.trade
//
// 【商旅】酒店交易查询流水接口——杭州市政府
type AlitripbtripopensupplychainhoteltradeAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenApiZzdSearchRq
}

// NewAlitripbtripopensupplychainhoteltradeRequest 初始化AlitripbtripopensupplychainhoteltradeAPIRequest对象
func NewAlitripbtripopensupplychainhoteltradeRequest() *AlitripbtripopensupplychainhoteltradeAPIRequest {
	return &AlitripbtripopensupplychainhoteltradeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripopensupplychainhoteltradeAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.supplychain.hotel.trade"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripopensupplychainhoteltradeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripopensupplychainhoteltradeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripbtripopensupplychainhoteltradeAPIRequest) SetRq(_rq *OpenApiZzdSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripopensupplychainhoteltradeAPIRequest) GetRq() *OpenApiZzdSearchRq {
	return r._rq
}
