package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembervouchercardpurchasablequery 德比付费会员卡可购查询
// alitrip.merchant.galaxy.derby.member.voucher.card.purchasable.query
//
// 德比付费会员卡可购查询
func Alitripmerchantgalaxyderbymembervouchercardpurchasablequery(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardpurchasablequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
