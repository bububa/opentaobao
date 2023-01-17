package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptShow v5.0付费会员卡开发发票图片展示
// alitrip.merchant.galaxy.derby.member.voucher.receipt.show
//
// v5.0付费会员卡开发发票图片展示
func AlitripMerchantGalaxyDerbyMemberVoucherReceiptShow(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherReceiptShowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
