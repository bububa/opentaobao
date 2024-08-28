package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyCommonBindMerchantId 绑定用户和merchantID
// alitrip.merchant.galaxy.common.bind.merchant.id
//
// 绑定用户和merchantID
func AlitripMerchantGalaxyCommonBindMerchantId(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyCommonBindMerchantIdAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyCommonBindMerchantIdAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
