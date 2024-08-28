package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyFavoriteList 用户收藏列表查询
// alitrip.merchant.galaxy.favorite.list
//
// 让用户可以查询自己收藏的酒店列表
func AlitripMerchantGalaxyFavoriteList(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyFavoriteListAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyFavoriteListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
