package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyvoucherquerylist 查询代金券列表
// alitrip.merchant.galaxy.voucher.query.list
//
// 查询代金券列表
func Alitripmerchantgalaxyvoucherquerylist(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyvoucherquerylistAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyvoucherquerylistAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyvoucherquerylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
