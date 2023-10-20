package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxymemberproviderregister 对外提供会员注册服务
// alitrip.merchant.galaxy.member.provider.register
//
// 星河产品=对外提供注册雅高会员服务
func Alitripmerchantgalaxymemberproviderregister(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxymemberproviderregisterAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxymemberproviderregisterAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxymemberproviderregisterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
