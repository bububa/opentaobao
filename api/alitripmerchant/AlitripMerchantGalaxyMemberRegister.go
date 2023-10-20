package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxymemberregister 星河-微信小程序会员注册
// alitrip.merchant.galaxy.member.register
//
// 星河产品=微信小程序注册雅高会员服务
func Alitripmerchantgalaxymemberregister(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxymemberregisterAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxymemberregisterAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxymemberregisterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
