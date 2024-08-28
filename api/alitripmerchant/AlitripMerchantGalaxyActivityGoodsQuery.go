package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyActivityGoodsQuery 营销抽奖-用户奖品查询
// alitrip.merchant.galaxy.activity.goods.query
//
// 星河产品-提供营销抽奖奖品查询服务
func AlitripMerchantGalaxyActivityGoodsQuery(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyActivityGoodsQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyActivityGoodsQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
