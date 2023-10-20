package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQuery v5.0付费会员卡开发订单开票信息查询
// alitrip.merchant.galaxy.derby.member.voucher.receipt.details.query
//
// v5.0付费会员卡开发订单开票信息查询
func AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherReceiptDetailsQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
