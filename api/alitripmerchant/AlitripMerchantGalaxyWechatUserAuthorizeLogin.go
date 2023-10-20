package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxywechatuserauthorizelogin DFC-ID用户手机号授权登录
// alitrip.merchant.galaxy.wechat.user.authorize.login
//
// DFC-ID用户手机号授权登录
func Alitripmerchantgalaxywechatuserauthorizelogin(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxywechatuserauthorizeloginAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxywechatuserauthorizeloginAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxywechatuserauthorizeloginAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
