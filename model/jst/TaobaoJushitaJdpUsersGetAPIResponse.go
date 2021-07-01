package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJushitaJdpUsersGetAPIResponse
获取开通的订单同步服务的用户 API返回值
taobao.jushita.jdp.users.get

获取开通的订单同步服务的用户，含有rds的路由关系 */
type TaobaoJushitaJdpUsersGetAPIResponse struct {
	model.CommonResponse
	TaobaoJushitaJdpUsersGetAPIResponseModel
}

// TaobaoJushitaJdpUsersGetAPIResponseModel is 获取开通的订单同步服务的用户 成功返回结果
type TaobaoJushitaJdpUsersGetAPIResponseModel struct {
	XMLName xml.Name `xml:"jushita_jdp_users_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户列表
	Users []JdpUser `json:"users,omitempty" xml:"users>jdp_user,omitempty"`
	// 总记录数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
