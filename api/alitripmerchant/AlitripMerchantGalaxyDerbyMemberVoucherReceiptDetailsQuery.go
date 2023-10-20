package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembervoucherreceiptdetailsquery v5.0付费会员卡开发订单开票信息查询
// alitrip.merchant.galaxy.derby.member.voucher.receipt.details.query
//
// v5.0付费会员卡开发订单开票信息查询
func Alitripmerchantgalaxyderbymembervoucherreceiptdetailsquery(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembervoucherreceiptdetailsqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
