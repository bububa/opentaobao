package alitripmerchant

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alitripmerchant"
)

/* AlitripMerchantGalaxyMemberQuery
星河-获取登录用户的信息
alitrip.merchant.galaxy.member.query

获取登录用户的信息 */
func AlitripMerchantGalaxyMemberQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberQueryAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyMemberQueryAPIResponse, error) {
	var resp alitripmerchant.AlitripMerchantGalaxyMemberQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
