package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembervouchercardredeem 根据兑换码兑换臻享卡接口
// alitrip.merchant.galaxy.derby.member.voucher.card.redeem
//
// 根据兑换码兑换臻享卡接口
func Alitripmerchantgalaxyderbymembervouchercardredeem(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardredeemAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardredeemAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardredeemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
