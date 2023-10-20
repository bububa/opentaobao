package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyfavoritesave 用户添加或移除收藏接口
// alitrip.merchant.galaxy.favorite.save
//
// 用户可以点击收藏酒店，再次点击移除收藏的酒店
func Alitripmerchantgalaxyfavoritesave(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyfavoritesaveAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyfavoritesaveAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyfavoritesaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
