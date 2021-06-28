package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
添加用户 APIResponse
taobao.openim.users.add

导入用户
*/
type TaobaoOpenimUsersAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoOpenimUsersAddResponse `json:"openim_users_add_response,omitempty"` 
    TaobaoOpenimUsersAddResponse
}

/* model for simplify = false
type TaobaoOpenimUsersAddResponse struct {

    // 成功用户列表
    
    UidSucc  struct {
        String  []string `json:"string,omitempty"`
    } `json:"uid_succ,omitempty"`
    

    // 添加失败的用户id
    
    UidFail  struct {
        String  []string `json:"string,omitempty"`
    } `json:"uid_fail,omitempty"`
    

    // 添加帐号失败的具体信息
    
    FailMsg  struct {
        String  []string `json:"string,omitempty"`
    } `json:"fail_msg,omitempty"`
    

}
*/

type TaobaoOpenimUsersAddResponse struct {

    // 成功用户列表
    UidSucc   []string `json:"uid_succ,omitempty"`

    // 添加失败的用户id
    UidFail   []string `json:"uid_fail,omitempty"`

    // 添加帐号失败的具体信息
    FailMsg   []string `json:"fail_msg,omitempty"`

}
