package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加用户 APIResponse
taobao.openim.users.add

导入用户
*/
type TaobaoOpenimUsersAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"openim_users_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 成功用户列表
    
    UidSucc   []string `json:"uid_succ,omitempty" xml:"