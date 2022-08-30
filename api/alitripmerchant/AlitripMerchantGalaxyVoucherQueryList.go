package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyVoucherQueryList 查询代金券列表
// alitrip.merchant.galaxy.voucher.query.list
//
// 查询代金券列表
func AlitripMerchantGalaxyVoucherQueryList(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyVoucherQueryListAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyVoucherQueryListAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyVoucherQueryListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
