package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembervouchercardactive 权益卡订单激活
// alitrip.merchant.galaxy.derby.member.voucher.card.active
//
// 权益卡订单激活
func Alitripmerchantgalaxyderbymembervouchercardactive(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardactiveAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardactiveAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardactiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
