package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyfavoritelist 用户收藏列表查询
// alitrip.merchant.galaxy.favorite.list
//
// 让用户可以查询自己收藏的酒店列表
func Alitripmerchantgalaxyfavoritelist(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyfavoritelistAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyfavoritelistAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyfavoritelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
