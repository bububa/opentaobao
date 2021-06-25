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
    Response *TaobaoJushitaJdpUsersGetResponse `json:"taobao_jushita_jdp_users_get_response,omitempty"`
}

type TaobaoJushitaJdpUsersGetResponse struct {

    // 用户列表
    Users   []JdpUser `json:"users,omitempty"`

    // 总记录数
    TotalResults   int64 `json:"total_results,omitempty"`

}
