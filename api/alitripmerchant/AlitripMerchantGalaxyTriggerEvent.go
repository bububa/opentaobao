package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyTriggerEvent 抽奖活动自定义事件触发
// alitrip.merchant.galaxy.trigger.event
//
// 自定义事件触发
func AlitripMerchantGalaxyTriggerEvent(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyTriggerEventAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyTriggerEventAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyTriggerEventAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
