package alitripmerchant

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alitripmerchant"
)

/* 
提供会员查询接口 
alitrip.merchant.galaxy.provider.member.query

星河产品=提供会查询服务
*/
func AlitripMerchantGalaxyProviderMemberQuery(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyProviderMemberQueryAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyProviderMemberQueryAPIResponse, error) {
    var resp alitripmerchant.AlitripMerchantGalaxyProviderMemberQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
