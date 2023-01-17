package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptAuto 德比付费会员卡开票自匹配
// alitrip.merchant.galaxy.derby.member.voucher.receipt.auto
//
// 德比付费会员卡开票自匹配
func AlitripMerchantGalaxyDerbyMemberVoucherReceiptAuto(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherReceiptAutoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
