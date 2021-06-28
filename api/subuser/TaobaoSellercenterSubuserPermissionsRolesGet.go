package subuser

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/subuser"
)

/* 
查询指定的子账号的权限和角色信息 
taobao.sellercenter.subuser.permissions.roles.get

查询指定的子账号的被直接赋予的权限信息和角色信息。<br/>返回对象中包括直接赋予子账号的权限点信息、被赋予的角色以及角色的对应权限点信息。
*/
func TaobaoSellercenterSubuserPermissionsRolesGet(clt *core.SDKClient, req *subuser.TaobaoSellercenterSubuserPermissionsRolesGetRequest, session string) (*subuser.TaobaoSellercenterSubuserPermissionsRolesGetAPIResponse, error) {
    var resp subuser.TaobaoSellercenterSubuserPermissionsRolesGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
