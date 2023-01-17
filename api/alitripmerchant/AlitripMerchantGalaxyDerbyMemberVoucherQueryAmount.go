package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherQueryAmount 查询用户拥有的臻享卡数量
// alitrip.merchant.galaxy.derby.member.voucher.query.amount
//
// 查询用户拥有的臻享卡数量
func AlitripMerchantGalaxyDerbyMemberVoucherQueryAmount(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherQueryAmountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
