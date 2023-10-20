package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyCityLike 星河-酒店城市模糊查询
// alitrip.merchant.galaxy.city.like
//
// 根据城市模糊查询，雅高酒店所在城市的城市信息
func AlitripMerchantGalaxyCityLike(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyCityLikeAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyCityLikeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
