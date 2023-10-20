package subuser

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosellercenterrolesgetAPIResponse 获取指定卖家的角色列表 API返回值
// taobao.sellercenter.roles.get
//
// 获取指定卖家的角色列表，只能获取属于登陆者自己的信息。
type TaobaosellercenterrolesgetAPIResponse struct {
	model.CommonResponse
	TaobaosellercenterrolesgetAPIResponseModel
}

// TaobaosellercenterrolesgetAPIResponseModel is 获取指定卖家的角色列表 成功返回结果
type TaobaosellercenterrolesgetAPIResponseModel struct {
	XMLName xml.Name `xml:"sellercenter_roles_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 卖家子账号角色列表。&lt;br/&gt;返回对象为 role数据对象中的role_id，role_name，description，seller_id，create_time，modified_time。不包含permissions(权限点)
	Roles []Role `json:"roles,omitempty" xml:"roles>role,omitempty"`
}
