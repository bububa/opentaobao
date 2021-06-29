package alimember

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alimember"
)

/* 
校验商家身份 
alibaba.member.checkmerchant

校验商家身份
*/
func AlibabaMemberCheckmerchant(clt *core.SDKClient, req *alimember.AlibabaMemberCheckmerchantRequest, session string) (*alimember.AlibabaMemberCheckmerchantAPIResponse, error) {
    var resp alimember.AlibabaMemberCheckmerchantAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
