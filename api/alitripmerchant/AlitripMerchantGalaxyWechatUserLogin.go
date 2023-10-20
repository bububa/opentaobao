package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxywechatuserlogin 微信小程序用户登录
// alitrip.merchant.galaxy.wechat.user.login
//
// 微信小程序用户登录接口
func Alitripmerchantgalaxywechatuserlogin(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxywechatuserloginAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxywechatuserloginAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxywechatuserloginAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
