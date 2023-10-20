package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyActivityFatigue 营销抽奖-弹窗疲劳度控制
// alitrip.merchant.galaxy.activity.fatigue
//
// 星河产品-营销抽奖首页弹窗疲劳度控制服务
func AlitripMerchantGalaxyActivityFatigue(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyActivityFatigueAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyActivityFatigueAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyActivityFatigueAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
