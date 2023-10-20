package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// Alitripmerchantgalaxyuserrisk 查询微信账号的风险等级
// alitrip.merchant.galaxy.user.risk
//
// 【星河服务】查询微信账号的风险等级
func Alitripmerchantgalaxyuserrisk(clt *core.SDKClient, req *alitripmerchant.AlitripmerchantgalaxyuserriskAPIRequest, session string) (*alitripmerchant.AlitripmerchantgalaxyuserriskAPIResponse, error) {
	var resp alitripmerchant.AlitripmerchantgalaxyuserriskAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
