package subuser

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取指定卖家的角色列表 APIResponse
taobao.sellercenter.roles.get

获取指定卖家的角色列表，只能获取属于登陆者自己的信息。
*/
type TaobaoSellercenterRolesGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSellercenterRolesGetResponse `json:"sellercenter_roles_get_response,omitempty"` 
    TaobaoSellercenterRolesGetResponse
}

/* model for simplify = false
type TaobaoSellercenterRolesGetResponse struct {

    // 卖家子账号角色列表。<br/>返回对象为 role数据对象中的role_id，role_name，description，seller_id，create_time，modified_time。不包含permissions(权限点)
    
    Roles  struct {
        Role  []Role `json:"role,omitempty"`
    } `json:"roles,omitempty"`
    

}
*/

type TaobaoSellercenterRolesGetResponse struct {

    // 卖家子账号角色列表。<br/>返回对象为 role数据对象中的role_id，role_name，description，seller_id，create_time，modified_time。不包含permissions(权限点)
    Roles   []Role `json:"roles,omitempty"`

}
