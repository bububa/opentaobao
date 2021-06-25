package openim

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openim"
)

/* 
批量获取用户信息 
taobao.openim.users.get

批量获取用户信息
*/
func TaobaoOpenimUsersGet(clt *core.SDKClient, req *openim.TaobaoOpenimUsersGetRequest, session string) (*openim.TaobaoOpenimUsersGetResponse, error) {
    var resp openim.TaobaoOpenimUsersGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
