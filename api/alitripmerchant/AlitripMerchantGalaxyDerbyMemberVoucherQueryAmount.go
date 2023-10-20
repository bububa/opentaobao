package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherQueryAmount 查询用户拥有的臻享卡数量
// alitrip.merchant.galaxy.derby.member.voucher.query.amount
//
// 查询用户拥有的臻享卡数量
func AlitripMerchantGalaxyDerbyMemberVoucherQueryAmount(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
