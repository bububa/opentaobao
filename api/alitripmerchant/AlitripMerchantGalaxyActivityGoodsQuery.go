package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyActivityGoodsQuery 营销抽奖-用户奖品查询
// alitrip.merchant.galaxy.activity.goods.query
//
// 星河产品-提供营销抽奖奖品查询服务
func AlitripMerchantGalaxyActivityGoodsQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyActivityGoodsQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyActivityGoodsQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
