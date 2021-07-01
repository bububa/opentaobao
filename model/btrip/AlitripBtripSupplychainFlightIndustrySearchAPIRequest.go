package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripSupplychainFlightIndustrySearchAPIRequest
机票行业搜索接口 API请求
alitrip.btrip.supplychain.flight.industry.search

【商旅】机票行业搜索 */
type AlitripBtripSupplychainFlightIndustrySearchAPIRequest struct {
	model.Params
	// 入参
	_rq *FlightSearchRq
}

// New
