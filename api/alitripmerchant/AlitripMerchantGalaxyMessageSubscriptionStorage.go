package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMessageSubscriptionStorage 存储模版ID
// alitrip.merchant.galaxy.message.subscription.storage
//
// 消息订阅中的消息模版的存储
func AlitripMerchantGalaxyMessageSubscriptionStorage(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyMessageSubscriptionStorageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
