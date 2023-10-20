package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyverifysignature 微信账号生物认证
// alitrip.merchant.galaxy.verify.signature
//
// 在用户注册等场景下，如果账号风险等级过高，需要进行生物认证
func Alitripmerchantgalaxyverifysignature(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyverifysignatureAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyverifysignatureAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyverifysignatureAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
