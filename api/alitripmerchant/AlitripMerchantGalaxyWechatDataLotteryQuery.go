package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxywechatdatalotteryquery 抽奖用户名单查询接口
// alitrip.merchant.galaxy.wechat.data.lottery.query
//
// 抽奖用户名单查询接口
func Alitripmerchantgalaxywechatdatalotteryquery(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxywechatdatalotteryqueryAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxywechatdatalotteryqueryAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxywechatdatalotteryqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
