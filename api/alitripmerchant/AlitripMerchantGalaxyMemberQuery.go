package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyMemberQuery 星河-获取登录用户的信息
// alitrip.merchant.galaxy.member.query
//
// 获取登录用户的信息
func AlitripMerchantGalaxyMemberQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyMemberQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
