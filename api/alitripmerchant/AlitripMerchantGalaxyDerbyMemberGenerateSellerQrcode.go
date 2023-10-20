package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcode 生成臻享卡德比分销二维码
// alitrip.merchant.galaxy.derby.member.generate.seller.qrcode
//
// 生成臻享卡德比分销二维码
func AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcode(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
