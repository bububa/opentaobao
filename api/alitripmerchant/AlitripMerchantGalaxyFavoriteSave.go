package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyFavoriteSave 用户添加或移除收藏接口
// alitrip.merchant.galaxy.favorite.save
//
// 用户可以点击收藏酒店，再次点击移除收藏的酒店
func AlitripMerchantGalaxyFavoriteSave(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyFavoriteSaveAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyFavoriteSaveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
