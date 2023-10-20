package subuser

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSellercenterSubuserPermissionsRolesGetAPIResponse 查询指定的子账号的权限和角色信息 API返回值
// taobao.sellercenter.subuser.permissions.roles.get
//
// 查询指定的子账号的被直接赋予的权限信息和角色信息。&lt;br/&gt;返回对象中包括直接赋予子账号的权限点信息、被赋予的角色以及角色的对应权限点信息。
type TaobaoSellercenterSubuserPermissionsRolesGetAPIResponse struct {
	model.CommonResponse
	TaobaoSellercenterSubuserPermissionsRolesGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSellercenterSubuserPermissionsRolesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSellercenterSubuserPermissionsRolesGetAPIResponseModel).Reset()
}

// TaobaoSellercenterSubuserPermissionsRolesGetAPIResponseModel is 查询指定的子账号的权限和角色信息 成功返回结果
type TaobaoSellercenterSubuserPermissionsRolesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"sellercenter_subuser_permissions_roles_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子账号被所拥有的权限
	SubuserPermission *SubUserPermission `json:"subuser_permission,omitempty" xml:"subuser_permission,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSellercenterSubuserPermissionsRolesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.SubuserPermission = nil
}

var poolTaobaoSellercenterSubuserPermissionsRolesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSellercenterSubuserPermissionsRolesGetAPIResponse)
	},
}

// GetTaobaoSellercenterSubuserPermissionsRolesGetAPIResponse 从 sync.Pool 获取 TaobaoSellercenterSubuserPermissionsRolesGetAPIResponse
func GetTaobaoSellercenterSubuserPermissionsRolesGetAPIResponse() *TaobaoSellercenterSubuserPermissionsRolesGetAPIResponse {
	return poolTaobaoSellercenterSubuserPermissionsRolesGetAPIResponse.Get().(*TaobaoSellercenterSubuserPermissionsRolesGetAPIResponse)
}

// ReleaseTaobaoSellercenterSubuserPermissionsRolesGetAPIResponse 将 TaobaoSellercenterSubuserPermissionsRolesGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSellercenterSubuserPermissionsRolesGetAPIResponse(v *TaobaoSellercenterSubuserPermissionsRolesGetAPIResponse) {
	v.Reset()
	poolTaobaoSellercenterSubuserPermissionsRolesGetAPIResponse.Put(v)
}
