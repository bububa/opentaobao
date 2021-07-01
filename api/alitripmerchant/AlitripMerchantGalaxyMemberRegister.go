package alitripmerchant

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alitripmerchant"
)

/* 
星河-微信小程序会员注册 
alitrip.merchant.galaxy.member.register

星河产品=微信小程序注册雅高会员服务
*/
func AlitripMerchantGalaxyMemberRegister(clt *core.SDKClient, req *alitripmerchant.AlitripMerchantGalaxyMemberRegisterAPIRequest, session string) (*alitripmerchant.AlitripMerchantGalaxyMemberRegisterAPIResponse, error) {
    var resp alitripmerchant.AlitripMerchantGalaxyMemberRegisterAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
