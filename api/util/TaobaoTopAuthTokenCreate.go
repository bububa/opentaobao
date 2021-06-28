package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
获取Access Token 
taobao.top.auth.token.create

用户通过code换获取access_token，https only
*/
func TaobaoTopAuthTokenCreate(clt *core.SDKClient, req *util.TaobaoTopAuthTokenCreateRequest, session string) (*util.TaobaoTopAuthTokenCreateAPIResponse, error) {
    var resp util.TaobaoTopAuthTokenCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
