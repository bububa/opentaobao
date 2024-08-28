package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyWechatUserModifyPhone DFC-ID用户换绑手机号
// alitrip.merchant.galaxy.wechat.user.modify.phone
//
// DFC-ID用户换绑手机号
func AlitripMerchantGalaxyWechatUserModifyPhone(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyWechatUserModifyPhoneAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyWechatUserModifyPhoneAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
