package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripSupplychainBusIndustrySearch 汽车票行业搜索接口
// alitrip.btrip.supplychain.bus.industry.search
//
// 汽车票行业搜索接口
func AlitripBtripSupplychainBusIndustrySearch(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripSupplychainBusIndustrySearchAPIRequest, resp *btrip.AlitripBtripSupplychainBusIndustrySearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
