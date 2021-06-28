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
    // Response *TaobaoOpenimUsersDeleteResponse `json:"openim_users_delete_response,omitempty"` 
    TaobaoOpenimUsersDeleteResponse
}

/* model for simplify = false
type TaobaoOpenimUsersDeleteResponse struct {

    // 操作成功
    
    Result  struct {
        String  []string `json:"string,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoOpenimUsersDeleteResponse struct {

    // 操作成功
    Result   []string `json:"result,omitempty"`

}
