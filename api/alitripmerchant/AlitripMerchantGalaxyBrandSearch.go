package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyBrandSearch 星河-品牌搜索
// alitrip.merchant.galaxy.brand.search
//
// 星河服务=获取雅高品牌信息
func AlitripMerchantGalaxyBrandSearch(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyBrandSearchAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyBrandSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
