package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembervouchercardorderdetailsquery 订单详情查询接口
// alitrip.merchant.galaxy.derby.member.voucher.card.order.details.query
//
// 订单详情查询接口
func Alitripmerchantgalaxyderbymembervouchercardorderdetailsquery(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardorderdetailsqueryAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardorderdetailsqueryAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardorderdetailsqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
