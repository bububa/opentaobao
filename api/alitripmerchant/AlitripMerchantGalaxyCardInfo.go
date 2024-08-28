package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyCardInfo 获取会员体系
// alitrip.merchant.galaxy.card.info
//
// 星河=根据卡类型获取当前的会员体系
func AlitripMerchantGalaxyCardInfo(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyCardInfoAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyCardInfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
