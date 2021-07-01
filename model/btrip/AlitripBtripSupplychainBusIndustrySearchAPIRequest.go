package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripSupplychainBusIndustrySearchAPIRequest
汽车票行业搜索接口 API请求
alitrip.btrip.supplychain.bus.industry.search

汽车票行业搜索接口 */
type AlitripBtripSupplychainBusIndustrySearchAPIRequest struct {
	model.Params
	// 入参
	_rq *BusSearchRq
}

// NewAlitripBtripSupplychainBusIndustrySearchRequest 初始化AlitripBtripSupplychainBusIndustrySearchAPIRequest对象
func NewAlitripBtripSupplychainBusIndustrySearchRequest() *AlitripBtripSupplychainBusIndustrySearchAPIRequest {
	return &AlitripBtripSupplychainBusIndustrySearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripSupplychainBusIndustrySearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.supplychain.bus.industry.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripSupplychainBusIndustrySearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Rq Setter
// 入参
func (r *AlitripBtripSupplychainBusIndustrySearchAPIRequest) SetRq(_rq *BusSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// Get Rq Getter
func (r AlitripBtripSupplychainBusIndustrySearchAPIRequest) GetRq() *BusSearchRq {
	return r._rq
}
