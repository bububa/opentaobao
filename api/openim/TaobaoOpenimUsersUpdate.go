package openim

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openim"
)

/* 
批量更新用户信息 
taobao.openim.users.update

批量更新用户信息
*/
func TaobaoOpenimUsersUpdate(clt *core.SDKClient, req *openim.TaobaoOpenimUsersUpdateRequest, session string) (*openim.TaobaoOpenimUsersUpdateAPIResponse, error) {
    var resp openim.TaobaoOpenimUsersUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
