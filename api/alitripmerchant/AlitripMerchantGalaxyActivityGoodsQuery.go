package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyactivitygoodsquery 营销抽奖-用户奖品查询
// alitrip.merchant.galaxy.activity.goods.query
//
// 星河产品-提供营销抽奖奖品查询服务
func Alitripmerchantgalaxyactivitygoodsquery(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyactivitygoodsqueryAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyactivitygoodsqueryAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyactivitygoodsqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
