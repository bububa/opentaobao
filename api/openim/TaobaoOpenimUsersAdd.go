package openim

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openim"
)

/* 
添加用户 
taobao.openim.users.add

导入用户
*/
func TaobaoOpenimUsersAdd(clt *core.SDKClient, req *openim.TaobaoOpenimUsersAddRequest, session string) (*openim.TaobaoOpenimUsersAddAPIResponse, error) {
    var resp openim.TaobaoOpenimUsersAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
