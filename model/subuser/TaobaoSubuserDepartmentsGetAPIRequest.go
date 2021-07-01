package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSubuserDepartmentsGetAPIRequest
获取指定账户的所有部门列表 API请求
taobao.subuser.departments.get

获取指定账户的所有部门列表，其实包括有每个部门的ID、父部门ID、部门名称（通过主账号登陆只能查询属于该主账号下的所有部门信息）。 */
type TaobaoSubuserDepartmentsGetAPIRequest struct {
	model.Params
	// 主账号用户名
	_userNick string
}

// New
