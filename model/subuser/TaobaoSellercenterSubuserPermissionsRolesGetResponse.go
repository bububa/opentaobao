package subuser

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询指定的子账号的权限和角色信息 APIResponse
taobao.sellercenter.subuser.permissions.roles.get

查询指定的子账号的被直接赋予的权限信息和角色信息。<br/>返回对象中包括直接赋予子账号的权限点信息、被赋予的角色以及角色的对应权限点信息。
*/
type TaobaoSellercenterSubuserPermissionsRolesGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSellercenterSubuserPermissionsRolesGetResponse `json:"sellercenter_subuser_permissions_roles_get_response,omitempty"` 
    TaobaoSellercenterSubuserPermissionsRolesGetResponse
}

/* model for simplify = false
type TaobaoSellercenterSubuserPermissionsRolesGetResponse struct {

    // 子账号被所拥有的权限
    
    SubuserPermission  *struct {
        SubUserPermission  *SubUserPermission `json:"sub_user_permission,omitempty"`
    } `json:"subuser_permission,omitempty"`
    

}
*/

type TaobaoSellercenterSubuserPermissionsRolesGetResponse struct {

    // 子账号被所拥有的权限
    SubuserPermission   *SubUserPermission `json:"subuser_permission,omitempty"`

}
