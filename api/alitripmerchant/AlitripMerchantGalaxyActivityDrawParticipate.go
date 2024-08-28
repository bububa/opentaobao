package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyActivityDrawParticipate 星河-幸运抽奖活动参与
// alitrip.merchant.galaxy.activity.draw.participate
//
// 雅高小程序幸运抽奖活动页面用户进行抽奖，根据活动规则返回抽奖结果
func AlitripMerchantGalaxyActivityDrawParticipate(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyActivityDrawParticipateAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyActivityDrawParticipateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
