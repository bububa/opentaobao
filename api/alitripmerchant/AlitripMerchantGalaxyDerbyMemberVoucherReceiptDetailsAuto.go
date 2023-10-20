package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembervoucherreceiptdetailsauto v5.0付费会员卡开票抬头自匹配
// alitrip.merchant.galaxy.derby.member.voucher.receipt.details.auto
//
// v5.0付费会员卡开票抬头自匹配
func Alitripmerchantgalaxyderbymembervoucherreceiptdetailsauto(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembervoucherreceiptdetailsautoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
