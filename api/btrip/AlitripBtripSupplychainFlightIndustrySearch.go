package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripSupplychainFlightIndustrySearch 机票行业搜索接口
// alitrip.btrip.supplychain.flight.industry.search
//
// 【商旅】机票行业搜索
func AlitripBtripSupplychainFlightIndustrySearch(clt *core.SDKClient, req *btrip.AlitripBtripSupplychainFlightIndustrySearchAPIRequest, resp *btrip.AlitripBtripSupplychainFlightIndustrySearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
