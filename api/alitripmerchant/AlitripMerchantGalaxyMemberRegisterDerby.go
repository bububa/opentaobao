package alitripmerchant

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMemberRegisterDerby 会员注册(新版注册接口对接德比)
// alitrip.merchant.galaxy.member.register.derby
//
// 会员注册(新版注册接口对接德比)，返回手机号/邮箱/注册状态
func AlitripMerchantGalaxyMemberRegisterDerby(ctx context.Context, clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberRegisterDerbyAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyMemberRegisterDerbyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
