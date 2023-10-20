package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyCommonBindMerchantId 绑定用户和merchantID
// alitrip.merchant.galaxy.common.bind.merchant.id
//
// 绑定用户和merchantID
func AlitripMerchantGalaxyCommonBindMerchantId(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyCommonBindMerchantIdAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyCommonBindMerchantIdAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
