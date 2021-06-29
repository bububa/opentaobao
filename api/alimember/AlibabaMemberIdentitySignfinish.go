package alimember

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alimember"
)

/* 
签约确认 
alibaba.member.identity.signfinish

签约确认
*/
func AlibabaMemberIdentitySignfinish(clt *core.SDKClient, req *alimember.AlibabaMemberIdentitySignfinishRequest, session string) (*alimember.AlibabaMemberIdentitySignfinishAPIResponse, error) {
    var resp alimember.AlibabaMemberIdentitySignfinishAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
