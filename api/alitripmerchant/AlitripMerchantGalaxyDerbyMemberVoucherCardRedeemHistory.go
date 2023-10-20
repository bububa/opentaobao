package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembervouchercardredeemhistory 查询会员兑换臻享卡历史记录
// alitrip.merchant.galaxy.derby.member.voucher.card.redeem.history
//
// 查询会员兑换臻享卡历史记录
func Alitripmerchantgalaxyderbymembervouchercardredeemhistory(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardredeemhistoryAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardredeemhistoryAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardredeemhistoryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
