package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardRedeem 根据兑换码兑换臻享卡接口
// alitrip.merchant.galaxy.derby.member.voucher.card.redeem
//
// 根据兑换码兑换臻享卡接口
func AlitripMerchantGalaxyDerbyMemberVoucherCardRedeem(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
