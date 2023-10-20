package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyFavoriteQuery 单酒店收藏状态查询
// alitrip.merchant.galaxy.favorite.query
//
// 返回用户对单个酒店的收藏状态
func AlitripMerchantGalaxyFavoriteQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyFavoriteQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyFavoriteQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
