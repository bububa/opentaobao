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
    Response *TaobaoOpenimUsersAddResponse `json:"taobao_openim_users_add_response,omitempty"`
}

type TaobaoOpenimUsersAddResponse struct {

    // 成功用户列表
    UidSucc   []String `json:"uid_succ,omitempty"`

    // 添加失败的用户id
    UidFail   []String `json:"uid_fail,omitempty"`

    // 添加帐号失败的具体信息
    FailMsg   []String `json:"fail_msg,omitempty"`

}
