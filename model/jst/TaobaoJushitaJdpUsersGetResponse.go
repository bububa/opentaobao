package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取开通的订单同步服务的用户 APIResponse
taobao.jushita.jdp.users.get

获取开通的订单同步服务的用户，含有rds的路由关系
*/
type TaobaoJushitaJdpUsersGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoJushitaJdpUsersGetResponse `json:"jushita_jdp_users_get_response,omitempty"` 
    TaobaoJushitaJdpUsersGetResponse
}

/* model for simplify = false
type TaobaoJushitaJdpUsersGetResponse struct {

    // 用户列表
    
    Users  struct {
        JdpUser  []JdpUser `json:"jdp_user,omitempty"`
    } `json:"users,omitempty"`
    

    // 总记录数
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

}
*/

type TaobaoJushitaJdpUsersGetResponse struct {

    // 用户列表
    Users   []JdpUser `json:"users,omitempty"`

    // 总记录数
    TotalResults   int64 `json:"total_results,omitempty"`

}
