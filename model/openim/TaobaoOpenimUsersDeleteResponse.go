package openim

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除用户 APIResponse
taobao.openim.users.delete

批量删除用户
*/
type TaobaoOpenimUsersDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoOpenimUsersDeleteResponse `json:"taobao_openim_users_delete_response,omitempty"`
}

type TaobaoOpenimUsersDeleteResponse struct {

    // 操作成功
    Result   []String `json:"result,omitempty"`

}
