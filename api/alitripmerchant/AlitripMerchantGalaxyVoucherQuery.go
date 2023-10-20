package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyVoucherQuery 查询单个代金券信息
// alitrip.merchant.galaxy.voucher.query
//
// 查询单个代金券信息
func AlitripMerchantGalaxyVoucherQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyVoucherQueryAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyVoucherQueryAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyVoucherQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
