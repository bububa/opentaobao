package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcode 会员权益卡身份识别二维码图片
// alitrip.merchant.galaxy.derby.member.voucher.card.show.qrcode
//
// 会员权益卡身份识别二维码图片
func AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcode(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
