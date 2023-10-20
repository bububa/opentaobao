package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembervouchercardquery 德比付费会员卡查询
// alitrip.merchant.galaxy.derby.member.voucher.card.query
//
// 德比付费会员卡查询
func Alitripmerchantgalaxyderbymembervouchercardquery(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardqueryAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardqueryAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembervouchercardqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
