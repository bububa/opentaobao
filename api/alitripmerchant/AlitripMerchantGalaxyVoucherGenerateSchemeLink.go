package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyVoucherGenerateSchemeLink 生成短信链接
// alitrip.merchant.galaxy.voucher.generate.scheme.link
//
// 生成微信跳转链接scheme_link
func AlitripMerchantGalaxyVoucherGenerateSchemeLink(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
