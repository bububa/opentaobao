package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatus 前端订单支付成功回调-修改订单状态
// alitrip.merchant.galaxy.derby.member.voucher.update.status
//
// 前端订单支付成功回调-修改订单状态
func AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatus(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherUpdateStatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
