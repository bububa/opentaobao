package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcode 生成臻享卡德比分销二维码
// alitrip.merchant.galaxy.derby.member.generate.seller.qrcode
//
// 生成臻享卡德比分销二维码
func AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcode(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
