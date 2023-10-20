package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembervouchercardorderplace 德比付费会员卡下单
// alitrip.merchant.galaxy.derby.member.voucher.card.order.place
//
// 德比付费会员卡下单
func Alitripmerchantgalaxyderbymembervouchercardorderplace(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardorderplaceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
