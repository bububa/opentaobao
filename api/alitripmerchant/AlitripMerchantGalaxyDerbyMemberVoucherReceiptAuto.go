package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembervoucherreceiptauto 德比付费会员卡开票自匹配
// alitrip.merchant.galaxy.derby.member.voucher.receipt.auto
//
// 德比付费会员卡开票自匹配
func Alitripmerchantgalaxyderbymembervoucherreceiptauto(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembervoucherreceiptautoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
