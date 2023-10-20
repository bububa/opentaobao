package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardActive 权益卡订单激活
// alitrip.merchant.galaxy.derby.member.voucher.card.active
//
// 权益卡订单激活
func AlitripMerchantGalaxyDerbyMemberVoucherCardActive(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardActiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
