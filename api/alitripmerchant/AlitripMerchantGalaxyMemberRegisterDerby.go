package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMemberRegisterDerby 会员注册(新版注册接口对接德比)
// alitrip.merchant.galaxy.member.register.derby
//
// 会员注册(新版注册接口对接德比)，返回手机号/邮箱/注册状态
func AlitripMerchantGalaxyMemberRegisterDerby(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberRegisterDerbyAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyMemberRegisterDerbyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
