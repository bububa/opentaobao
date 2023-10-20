package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxymemberlogout 星河-用户登出
// alitrip.merchant.galaxy.member.logout
//
// 星河=微信小程序用户登出
func Alitripmerchantgalaxymemberlogout(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxymemberlogoutAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxymemberlogoutAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxymemberlogoutAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
