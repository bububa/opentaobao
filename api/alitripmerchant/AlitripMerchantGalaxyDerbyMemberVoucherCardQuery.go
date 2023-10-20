package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardQuery 德比付费会员卡查询
// alitrip.merchant.galaxy.derby.member.voucher.card.query
//
// 德比付费会员卡查询
func AlitripMerchantGalaxyDerbyMemberVoucherCardQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
