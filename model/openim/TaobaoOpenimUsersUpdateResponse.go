package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
批量更新用户信息 APIResponse
taobao.openim.users.update

批量更新用户信息
*/
type TaobaoOpenimUsersUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenimUsersUpdateResponse `json:"openim_users_update_response,omitempty"` 
    TaobaoOpenimUsersUpdateResponse
}

/* model for simplify = false
type TaobaoOpenimUsersUpdateResponse struct {

    // 对应每一个失败用户的具体错误信息
    
    FailMsg  struct {
        String  []string `json:"string,omitempty"`
    } `json:"fail_msg,omitempty"`
    

    // 失败的uid列表
    
    UidFail  struct {
        String  []string `json:"string,omitempty"`
    } `json:"uid_fail,omitempty"`
    

    // 成功的uid列表
    
    UidSucc  struct {
        String  []string `json:"string,omitempty"`
    } `json:"uid_succ,omitempty"`
    

}
*/

type TaobaoOpenimUsersUpdateResponse struct {

    // 对应每一个失败用户的具体错误信息
    FailMsg   []string `json:"fail_msg,omitempty"`

    // 失败的uid列表
    UidFail   []string `json:"uid_fail,omitempty"`

    // 成功的uid列表
    UidSucc   []string `json:"uid_succ,omitempty"`

}
