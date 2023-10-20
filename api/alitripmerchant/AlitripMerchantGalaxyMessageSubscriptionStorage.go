package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMessageSubscriptionStorage 存储模版ID
// alitrip.merchant.galaxy.message.subscription.storage
//
// 消息订阅中的消息模版的存储
func AlitripMerchantGalaxyMessageSubscriptionStorage(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMessageSubscriptionStorageAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyMessageSubscriptionStorageAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyMessageSubscriptionStorageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
