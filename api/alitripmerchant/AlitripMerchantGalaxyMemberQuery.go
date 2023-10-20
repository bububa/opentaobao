package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxymemberquery 星河-获取登录用户的信息
// alitrip.merchant.galaxy.member.query
//
// 获取登录用户的信息
func Alitripmerchantgalaxymemberquery(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxymemberqueryAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxymemberqueryAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxymemberqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
