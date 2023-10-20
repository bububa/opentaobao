package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancel 取消订单
// alitrip.merchant.galaxy.derby.member.voucher.card.order.cancel
//
// 德比付费会员卡订单取消
func AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancel(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardOrderCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
