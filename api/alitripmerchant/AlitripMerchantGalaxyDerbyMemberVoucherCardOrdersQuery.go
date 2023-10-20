package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembervouchercardordersquery 查询权益卡订单列表
// alitrip.merchant.galaxy.derby.member.voucher.card.orders.query
//
// 查询权益卡订单列表
func Alitripmerchantgalaxyderbymembervouchercardordersquery(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardordersqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
