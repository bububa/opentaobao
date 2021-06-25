package simba

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/simba"
)

/* 
查看功能权限 
taobao.simba.customers.sid.get

查询用户是否拥有某个功能权限
*/
func TaobaoSimbaCustomersSidGet(clt *core.SDKClient, req *simba.TaobaoSimbaCustomersSidGetRequest, session string) (*simba.TaobaoSimbaCustomersSidGetResponse, error) {
    var resp simba.TaobaoSimbaCustomersSidGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
