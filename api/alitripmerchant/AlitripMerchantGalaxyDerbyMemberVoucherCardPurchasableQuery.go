package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQuery 德比付费会员卡可购查询
// alitrip.merchant.galaxy.derby.member.voucher.card.purchasable.query
//
// 德比付费会员卡可购查询
func AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
