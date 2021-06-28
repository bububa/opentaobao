package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量更新用户信息 APIResponse
taobao.openim.users.update

批量更新用户信息
*/
type TaobaoOpenimUsersUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"openim_users_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 对应每一个失败用户的具体错误信息
    
    FailMsg   []string `json:"fail_msg,omitempty" xml:"