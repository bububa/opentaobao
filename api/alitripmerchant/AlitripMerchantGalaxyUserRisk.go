package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyUserRisk 查询微信账号的风险等级
// alitrip.merchant.galaxy.user.risk
//
// 【星河服务】查询微信账号的风险等级
func AlitripMerchantGalaxyUserRisk(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyUserRiskAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyUserRiskAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyUserRiskAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
