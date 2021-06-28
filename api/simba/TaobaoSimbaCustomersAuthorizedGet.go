package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
取得当前登录用户的授权账户列表 
taobao.simba.customers.authorized.get

取得当前登录用户的授权账户列表
*/
func TaobaoSimbaCustomersAuthorizedGet(clt *core.SDKClient, req *simba.TaobaoSimbaCustomersAuthorizedGetRequest, session string) (*simba.TaobaoSimbaCustomersAuthorizedGetAPIResponse, error) {
    var resp simba.TaobaoSimbaCustomersAuthorizedGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
