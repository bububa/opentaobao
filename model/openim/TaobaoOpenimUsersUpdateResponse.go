package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量更新用户信息 API返回值 
taobao.openim.users.update

批量更新用户信息
*/
type TaobaoOpenimUsersUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimUsersUpdateResponse
}

// 批量更新用户信息 成功返回结果
type TaobaoOpenimUsersUpdateResponse struct {
    XMLName xml.Name `xml:"openim_users_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 对应每一个失败用户的具体错误信息
    FailMsg   []string `json:"fail_msg,omitempty" xml:"fail_msg>string,omitempty"`
    // 失败的uid列表
    UidFail   []string `json:"uid_fail,omitempty" xml:"uid_fail>string,omitempty"`
    // 成功的uid列表
    UidSucc   []string `json:"uid_succ,omitempty" xml:"uid_succ>string,omitempty"`
}
