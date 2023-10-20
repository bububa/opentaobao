package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyfavoritequery 单酒店收藏状态查询
// alitrip.merchant.galaxy.favorite.query
//
// 返回用户对单个酒店的收藏状态
func Alitripmerchantgalaxyfavoritequery(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyfavoritequeryAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyfavoritequeryAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyfavoritequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
