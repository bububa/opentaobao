package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyCommonBindMerchantId 绑定用户和merchantID
// alitrip.merchant.galaxy.common.bind.merchant.id
//
// 绑定用户和merchantID
func AlitripMerchantGalaxyCommonBindMerchantId(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyCommonBindMerchantIdAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyCommonBindMerchantIdAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyCommonBindMerchantIdAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
