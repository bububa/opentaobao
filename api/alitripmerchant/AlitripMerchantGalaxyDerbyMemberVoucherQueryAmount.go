package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembervoucherqueryamount 查询用户拥有的臻享卡数量
// alitrip.merchant.galaxy.derby.member.voucher.query.amount
//
// 查询用户拥有的臻享卡数量
func Alitripmerchantgalaxyderbymembervoucherqueryamount(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembervoucherqueryamountAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembervoucherqueryamountAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembervoucherqueryamountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
