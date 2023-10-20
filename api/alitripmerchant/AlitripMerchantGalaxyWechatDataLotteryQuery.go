package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatDataLotteryQuery 抽奖用户名单查询接口
// alitrip.merchant.galaxy.wechat.data.lottery.query
//
// 抽奖用户名单查询接口
func AlitripMerchantGalaxyWechatDataLotteryQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyWechatDataLotteryQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
