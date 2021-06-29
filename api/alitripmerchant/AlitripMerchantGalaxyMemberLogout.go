package alitripmerchant

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alitripmerchant"
)

/* 
星河-用户登出 
alitrip.merchant.galaxy.member.logout

星河=微信小程序用户登出
*/
func AlitripMerchantGalaxyMemberLogout(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberLogoutRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyMemberLogoutAPIResponse, error) {
    var resp alitripmerchant.AlitripMerchantGalaxyMemberLogoutAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
