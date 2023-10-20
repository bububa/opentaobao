package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripSupplychainBusIndustrySearch 汽车票行业搜索接口
// alitrip.btrip.supplychain.bus.industry.search
//
// 汽车票行业搜索接口
func AlitripBtripSupplychainBusIndustrySearch(clt *core.SDKClient, req *btrip.AlitripBtripSupplychainBusIndustrySearchAPIRequest, resp *btrip.AlitripBtripSupplychainBusIndustrySearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
