package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembervoucherreceiptdetailsapply v5.0付费会员卡开发订单开票详情申请
// alitrip.merchant.galaxy.derby.member.voucher.receipt.details.apply
//
// v5.0付费会员卡开发订单开票详情申请
func Alitripmerchantgalaxyderbymembervoucherreceiptdetailsapply(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembervoucherreceiptdetailsapplyAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembervoucherreceiptdetailsapplyAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembervoucherreceiptdetailsapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
