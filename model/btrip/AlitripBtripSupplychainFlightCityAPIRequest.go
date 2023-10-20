package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripsupplychainflightcityAPIRequest 机场数据查询 API请求
// alitrip.btrip.supplychain.flight.city
//
// 机场数据查询
type AlitripbtripsupplychainflightcityAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenSuggestRq
}

// NewAlitripbtripsupplychainflightcityRequest 初始化AlitripbtripsupplychainflightcityAPIRequest对象
func NewAlitripbtripsupplychainflightcityRequest() *AlitripbtripsupplychainflightcityAPIRequest {
	return &AlitripbtripsupplychainflightcityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripsupplychainflightcityAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.flight.city"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripsupplychainflightcityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripsupplychainflightcityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripsupplychainflightcityAPIRequest) SetRq(_rq *OpenSuggestRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripsupplychainflightcityAPIRequest) GetRq() *OpenSuggestRq {
	return r._rq
}
