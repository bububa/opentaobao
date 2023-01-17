package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQuery 订单详情查询接口
// alitrip.merchant.galaxy.derby.member.voucher.card.order.details.query
//
// 订单详情查询接口
func AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardOrderDetailsQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
