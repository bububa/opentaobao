package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyderbymembergeneratesellerqrcode 生成臻享卡德比分销二维码
// alitrip.merchant.galaxy.derby.member.generate.seller.qrcode
//
// 生成臻享卡德比分销二维码
func Alitripmerchantgalaxyderbymembergeneratesellerqrcode(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyderbymembergeneratesellerqrcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
