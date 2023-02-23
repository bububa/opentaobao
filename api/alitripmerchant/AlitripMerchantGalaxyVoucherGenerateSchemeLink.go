package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyVoucherGenerateSchemeLink 生成短信链接
// alitrip.merchant.galaxy.voucher.generate.scheme.link
//
// 生成微信跳转链接scheme_link
func AlitripMerchantGalaxyVoucherGenerateSchemeLink(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyVoucherGenerateSchemeLinkAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
