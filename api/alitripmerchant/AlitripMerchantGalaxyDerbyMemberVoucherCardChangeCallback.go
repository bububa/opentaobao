package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembervouchercardchangecallback v5.0德比付费会员卡通知
// alitrip.merchant.galaxy.derby.member.voucher.card.change.callback
//
// v5.0德比付费会员卡通知
func Alitripmerchantgalaxyderbymembervouchercardchangecallback(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardchangecallbackAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardchangecallbackAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardchangecallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
