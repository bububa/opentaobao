package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyOrderFill 填单页接口
// alitrip.merchant.galaxy.order.fill
//
// 进入填单页时调用此接口，返回填单所需展示基础信息
func AlitripMerchantGalaxyOrderFill(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyOrderFillAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyOrderFillAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
