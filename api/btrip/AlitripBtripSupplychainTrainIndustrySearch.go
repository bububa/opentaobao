package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripSupplychainTrainIndustrySearch 火车票行业搜索接口
// alitrip.btrip.supplychain.train.industry.search
//
// 【商旅】火车票行业搜索接口
func AlitripBtripSupplychainTrainIndustrySearch(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripSupplychainTrainIndustrySearchAPIRequest, resp *btrip.AlitripBtripSupplychainTrainIndustrySearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
