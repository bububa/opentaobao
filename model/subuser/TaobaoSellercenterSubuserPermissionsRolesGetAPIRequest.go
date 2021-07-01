package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSellercenterSubuserPermissionsRolesGetAPIRequest
查询指定的子账号的权限和角色信息 API请求
taobao.sellercenter.subuser.permissions.roles.get

查询指定的子账号的被直接赋予的权限信息和角色信息。<br/>返回对象中包括直接赋予子账号的权限点信息、被赋予的角色以及角色的对应权限点信息。 */
type TaobaoSellercenterSubuserPermissionsRolesGetAPIRequest struct {
	model.Params
	// 子账号昵称(子账号标识)
	_nick string
}

// New
