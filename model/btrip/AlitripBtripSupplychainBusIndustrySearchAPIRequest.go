package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripsupplychainbusindustrysearchAPIRequest 汽车票行业搜索接口 API请求
// alitrip.btrip.supplychain.bus.industry.search
//
// 汽车票行业搜索接口
type AlitripbtripsupplychainbusindustrysearchAPIRequest struct {
	model.Params
	// 入参
	_rq *BusSearchRq
}

// NewAlitripbtripsupplychainbusindustrysearchRequest 初始化AlitripbtripsupplychainbusindustrysearchAPIRequest对象
func NewAlitripbtripsupplychainbusindustrysearchRequest() *AlitripbtripsupplychainbusindustrysearchAPIRequest {
	return &AlitripbtripsupplychainbusindustrysearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripsupplychainbusindustrysearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.bus.industry.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripsupplychainbusindustrysearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripsupplychainbusindustrysearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripbtripsupplychainbusindustrysearchAPIRequest) SetRq(_rq *BusSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripsupplychainbusindustrysearchAPIRequest) GetRq() *BusSearchRq {
	return r._rq
}
