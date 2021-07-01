package openim

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimUsersAddAPIResponse
添加用户 API返回值
taobao.openim.users.add

导入用户 */
type TaobaoOpenimUsersAddAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimUsersAddAPIResponseModel
}

// TaobaoOpenimUsersAddAPIResponseModel is 添加用户 成功返回结果
type TaobaoOpenimUsersAddAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_users_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功用户列表
	UidSucc []string `json:"uid_succ,omitempty" xml:"uid_succ>string,omitempty"`
	// 添加失败的用户id
	UidFail []string `json:"uid_fail,omitempty" xml:"uid_fail>string,omitempty"`
	// 添加帐号失败的具体信息
	FailMsg []string `json:"fail_msg,omitempty" xml:"fail_msg>string,omitempty"`
}
