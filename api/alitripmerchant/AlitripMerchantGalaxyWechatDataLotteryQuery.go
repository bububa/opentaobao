package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatDataLotteryQuery 抽奖用户名单查询接口
// alitrip.merchant.galaxy.wechat.data.lottery.query
//
// 抽奖用户名单查询接口
func AlitripMerchantGalaxyWechatDataLotteryQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyWechatDataLotteryQueryAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyWechatDataLotteryQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
