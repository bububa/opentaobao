package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
获取登陆权限签名 
taobao.simba.login.authsign.get

获取登陆权限签名
*/
func TaobaoSimbaLoginAuthsignGet(clt *core.SDKClient, req *simba.TaobaoSimbaLoginAuthsignGetRequest, session string) (*simba.TaobaoSimbaLoginAuthsignGetAPIResponse, error) {
    var resp simba.TaobaoSimbaLoginAuthsignGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
