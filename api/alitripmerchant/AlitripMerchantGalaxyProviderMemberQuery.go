package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

// AlitripMerchantGalaxyProviderMemberQuery 提供会员查询接口
// alitrip.merchant.galaxy.provider.member.query
//
// 星河产品=提供会查询服务
func AlitripMerchantGalaxyProviderMemberQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyProviderMemberQueryAPIRequest, resp *alitripmerchant.AlitripMerchantGalaxyProviderMemberQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
