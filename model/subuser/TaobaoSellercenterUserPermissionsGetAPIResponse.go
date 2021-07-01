package subuser

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定用户的权限集合 API返回值 
taobao.sellercenter.user.permissions.get

获取指定用户的权限集合，并不组装成树。如果是主账号，返回所有的权限列表；如果是子账号，返回所有已授权的权限。只能查询属于自己的账号信息 (如果是主账号，则是主账号以及所属子账号，如果是子账号则是对应主账号以及所属子账号)
*/
type TaobaoSellercenterUserPermissionsGetAPIResponse struct {
    model.CommonResponse
    TaobaoSellercenterUserPermissionsGetAPIResponseModel
}

// 获取指定用户的权限集合 成功返回结果
type TaobaoSellercenterUserPermissionsGetAPIResponseModel struct {
    XMLName xml.Name `xml:"sellercenter_user_permissions_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 权限列表
    Permissions   []Permission `json:"permissions,omitempty" xml:"permissions>permission,omitempty"`
}
