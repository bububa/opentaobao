package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardActive 权益卡订单激活
// alitrip.merchant.galaxy.derby.member.voucher.card.active
//
// 权益卡订单激活
func AlitripMerchantGalaxyDerbyMemberVoucherCardActive(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
