package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyOfferQuery 星河-offer查询
// alitrip.merchant.galaxy.offer.query
//
// 根据offer的ID查询offer信息
func AlitripMerchantGalaxyOfferQuery(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyOfferQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyOfferQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
