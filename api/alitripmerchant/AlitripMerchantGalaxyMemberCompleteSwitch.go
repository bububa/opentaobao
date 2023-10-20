package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxymembercompleteswitch 会员切换模式接口
// alitrip.merchant.galaxy.member.complete.switch
//
// 小程序老用户调用德比接口进行会员切换
func Alitripmerchantgalaxymembercompleteswitch(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxymembercompleteswitchAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxymembercompleteswitchAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxymembercompleteswitchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
