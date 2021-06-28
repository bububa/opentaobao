package openim

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openim"
)

/* 
删除用户 
taobao.openim.users.delete

批量删除用户
*/
func TaobaoOpenimUsersDelete(clt *core.SDKClient, req *openim.TaobaoOpenimUsersDeleteRequest, session string) (*openim.TaobaoOpenimUsersDeleteAPIResponse, error) {
    var resp openim.TaobaoOpenimUsersDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
