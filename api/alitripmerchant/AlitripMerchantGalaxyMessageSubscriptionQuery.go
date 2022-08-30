package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMessageSubscriptionQuery 查询用户是否有模版ID权限
// alitrip.merchant.galaxy.message.subscription.query
//
// 只是查询用户是否拥有权限ID
func AlitripMerchantGalaxyMessageSubscriptionQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMessageSubscriptionQueryAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyMessageSubscriptionQueryAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyMessageSubscriptionQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
