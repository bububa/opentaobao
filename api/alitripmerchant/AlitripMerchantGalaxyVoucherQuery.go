package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyVoucherQuery 查询单个代金券信息
// alitrip.merchant.galaxy.voucher.query
//
// 查询单个代金券信息
func AlitripMerchantGalaxyVoucherQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyVoucherQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyVoucherQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
