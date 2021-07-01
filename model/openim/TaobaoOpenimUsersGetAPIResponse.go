package openim

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimUsersGetAPIResponse
批量获取用户信息 API返回值
taobao.openim.users.get

批量获取用户信息 */
type TaobaoOpenimUsersGetAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimUsersGetAPIResponseModel
}

// TaobaoOpenimUsersGetAPIResponseModel is 批量获取用户信息 成功返回结果
type TaobaoOpenimUsersGetAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_users_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 获取的用户信息列表
	Userinfos []Userinfos `json:"userinfos,omitempty" xml:"userinfos>userinfos,omitempty"`
}
