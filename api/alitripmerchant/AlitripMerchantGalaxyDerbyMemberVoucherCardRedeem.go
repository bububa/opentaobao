package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardRedeem 根据兑换码兑换臻享卡接口
// alitrip.merchant.galaxy.derby.member.voucher.card.redeem
//
// 根据兑换码兑换臻享卡接口
func AlitripMerchantGalaxyDerbyMemberVoucherCardRedeem(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
