package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefund Aps退券通知接口
// alitrip.merchant.galaxy.derby.member.voucher.card.aps.refund
//
// Aps退券通知接口
func AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefund(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
