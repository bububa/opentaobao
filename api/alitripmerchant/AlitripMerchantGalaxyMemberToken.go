package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxymembertoken 星河-校验token
// alitrip.merchant.galaxy.member.token
//
// 校验或者刷新token
func Alitripmerchantgalaxymembertoken(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxymembertokenAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxymembertokenAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxymembertokenAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
