package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefund Aps退券通知接口
// alitrip.merchant.galaxy.derby.member.voucher.card.aps.refund
//
// Aps退券通知接口
func AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefund(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardApsRefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
