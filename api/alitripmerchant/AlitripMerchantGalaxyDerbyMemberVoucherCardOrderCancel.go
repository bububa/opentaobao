package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembervouchercardordercancel 取消订单
// alitrip.merchant.galaxy.derby.member.voucher.card.order.cancel
//
// 德比付费会员卡订单取消
func Alitripmerchantgalaxyderbymembervouchercardordercancel(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardordercancelAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardordercancelAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardordercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
