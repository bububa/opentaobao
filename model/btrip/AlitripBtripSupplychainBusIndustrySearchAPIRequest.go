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

// New
