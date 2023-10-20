package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripsupplychainflightindustrysearchAPIRequest 机票行业搜索接口 API请求
// alitrip.btrip.supplychain.flight.industry.search
//
// 【商旅】机票行业搜索
type AlitripbtripsupplychainflightindustrysearchAPIRequest struct {
	model.Params
	// 入参
	_rq *FlightSearchRq
}

// NewAlitripbtripsupplychainflightindustrysearchRequest 初始化AlitripbtripsupplychainflightindustrysearchAPIRequest对象
func NewAlitripbtripsupplychainflightindustrysearchRequest() *AlitripbtripsupplychainflightindustrysearchAPIRequest {
	return &AlitripbtripsupplychainflightindustrysearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripsupplychainflightindustrysearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.flight.industry.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripsupplychainflightindustrysearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripsupplychainflightindustrysearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripbtripsupplychainflightindustrysearchAPIRequest) SetRq(_rq *FlightSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripsupplychainflightindustrysearchAPIRequest) GetRq() *FlightSearchRq {
	return r._rq
}
