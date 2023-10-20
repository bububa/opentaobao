package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyvoucherquery 查询单个代金券信息
// alitrip.merchant.galaxy.voucher.query
//
// 查询单个代金券信息
func Alitripmerchantgalaxyvoucherquery(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyvoucherqueryAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyvoucherqueryAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyvoucherqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
