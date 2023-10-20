package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMemberCompleteSwitch 会员切换模式接口
// alitrip.merchant.galaxy.member.complete.switch
//
// 小程序老用户调用德比接口进行会员切换
func AlitripMerchantGalaxyMemberCompleteSwitch(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberCompleteSwitchAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyMemberCompleteSwitchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
