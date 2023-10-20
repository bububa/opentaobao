package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembervouchercardapsrefund Aps退券通知接口
// alitrip.merchant.galaxy.derby.member.voucher.card.aps.refund
//
// Aps退券通知接口
func Alitripmerchantgalaxyderbymembervouchercardapsrefund(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardapsrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
