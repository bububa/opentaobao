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
    // Response *TaobaoOpenimUsersGetResponse `json:"openim_users_get_response,omitempty"` 
    TaobaoOpenimUsersGetResponse
}

/* model for simplify = false
type TaobaoOpenimUsersGetResponse struct {

    // 获取的用户信息列表
    
    Userinfos  struct {
        Userinfos  []Userinfos `json:"userinfos,omitempty"`
    } `json:"userinfos,omitempty"`
    

}
*/

type TaobaoOpenimUsersGetResponse struct {

    // 获取的用户信息列表
    Userinfos   []Userinfos `json:"userinfos,omitempty"`

}
