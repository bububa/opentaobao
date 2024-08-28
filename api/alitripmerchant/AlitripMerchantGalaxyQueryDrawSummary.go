package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyQueryDrawSummary 星河-抽奖活动概要列表查询
// alitrip.merchant.galaxy.query.draw.summary
//
// 雅高小程序抽奖活动列表概要查询
func AlitripMerchantGalaxyQueryDrawSummary(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyQueryDrawSummaryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyQueryDrawSummaryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
