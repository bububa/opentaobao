package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcode 生成臻享卡德比分销二维码
// alitrip.merchant.galaxy.derby.member.generate.seller.qrcode
//
// 生成臻享卡德比分销二维码
func AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcode(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
