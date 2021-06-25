package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量获取用户信息 APIResponse
taobao.openim.users.get

批量获取用户信息
*/
type TaobaoOpenimUsersGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOpenimUsersGetResponse `json:"taobao_openim_users_get_response,omitempty"`
}

type TaobaoOpenimUsersGetResponse struct {

    // 获取的用户信息列表
    Userinfos   []Userinfos `json:"userinfos,omitempty"`

}
