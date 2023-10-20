package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyvouchergenerateschemelink 生成短信链接
// alitrip.merchant.galaxy.voucher.generate.scheme.link
//
// 生成微信跳转链接scheme_link
func Alitripmerchantgalaxyvouchergenerateschemelink(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyvouchergenerateschemelinkAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyvouchergenerateschemelinkAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyvouchergenerateschemelinkAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
