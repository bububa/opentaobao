package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyTriggerEvent 抽奖活动自定义事件触发
// alitrip.merchant.galaxy.trigger.event
//
// 自定义事件触发
func AlitripMerchantGalaxyTriggerEvent(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyTriggerEventAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyTriggerEventAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
