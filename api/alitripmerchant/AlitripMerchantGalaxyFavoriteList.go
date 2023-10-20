package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyFavoriteList 用户收藏列表查询
// alitrip.merchant.galaxy.favorite.list
//
// 让用户可以查询自己收藏的酒店列表
func AlitripMerchantGalaxyFavoriteList(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyFavoriteListAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyFavoriteListAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyFavoriteListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
