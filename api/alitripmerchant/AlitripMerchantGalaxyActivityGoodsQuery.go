package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyActivityGoodsQuery 营销抽奖-用户奖品查询
// alitrip.merchant.galaxy.activity.goods.query
//
// 星河产品-提供营销抽奖奖品查询服务
func AlitripMerchantGalaxyActivityGoodsQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyActivityGoodsQueryAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyActivityGoodsQueryAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyActivityGoodsQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
