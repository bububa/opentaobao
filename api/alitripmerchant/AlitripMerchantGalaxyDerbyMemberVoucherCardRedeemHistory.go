package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistory 查询会员兑换臻享卡历史记录
// alitrip.merchant.galaxy.derby.member.voucher.card.redeem.history
//
// 查询会员兑换臻享卡历史记录
func AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistory(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardRedeemHistoryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
