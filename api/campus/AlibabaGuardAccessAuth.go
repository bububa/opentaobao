package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
鉴权 
alibaba.guard.access.auth

刷卡鉴权
*/
func AlibabaGuardAccessAuth(clt *core.SDKClient, req *campus.AlibabaGuardAccessAuthRequest, session string) (*campus.AlibabaGuardAccessAuthAPIResponse, error) {
    var resp campus.AlibabaGuardAccessAuthAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
