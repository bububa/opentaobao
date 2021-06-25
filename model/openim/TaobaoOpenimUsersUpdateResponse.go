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
    Response *TaobaoOpenimUsersUpdateResponse `json:"taobao_openim_users_update_response,omitempty"`
}

type TaobaoOpenimUsersUpdateResponse struct {

    // 对应每一个失败用户的具体错误信息
    FailMsg   []String `json:"fail_msg,omitempty"`

    // 失败的uid列表
    UidFail   []String `json:"uid_fail,omitempty"`

    // 成功的uid列表
    UidSucc   []String `json:"uid_succ,omitempty"`

}
