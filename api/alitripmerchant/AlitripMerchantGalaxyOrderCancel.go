package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyOrderCancel 星河-取消预订
// alitrip.merchant.galaxy.order.cancel
//
// 雅高酒店用户使用该接口，取消酒店预订
func AlitripMerchantGalaxyOrderCancel(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyOrderCancelAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyOrderCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
