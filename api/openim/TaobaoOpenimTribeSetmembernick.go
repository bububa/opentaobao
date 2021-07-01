package openim

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/openim"
)

/* 
设置群成员昵称 
taobao.openim.tribe.setmembernick

设置群成员昵称，存在如下两种场景
1 群主或管理员设置群成员昵称，该操作有权限控制。只针对普通群的群主和管理员开发此功能；讨论组群主不支持此设置操作
2 群成员设置自己的昵称，该功能对群所有成员开放
*/
func TaobaoOpenimTribeSetmembernick(clt *core.SDKClient, req *openim.TaobaoOpenimTribeSetmembernickAPIRequest, session string) (*openim.TaobaoOpenimTribeSetmembernickAPIResponse, error) {
    var resp openim.TaobaoOpenimTribeSetmembernickAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
