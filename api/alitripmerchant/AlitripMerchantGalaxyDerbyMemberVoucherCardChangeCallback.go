package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallback v5.0德比付费会员卡通知
// alitrip.merchant.galaxy.derby.member.voucher.card.change.callback
//
// v5.0德比付费会员卡通知
func AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallback(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardChangeCallbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
