package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclCheckemproleAPIRequest
校验用户是否有该角色 API请求
alibaba.campus.acl.checkemprole

校验用户是否有该权限 */
type AlibabaCampusAclCheckemproleAPIRequest struct {
	model.Params
	// 公司id不填默认为SYS_000
	_companyId int64
	// 系统id
	_systemId string
	// 园区id
	_campusId int64
	// 员工账号
	_accountId string
	// 角色id
	_itemKey string
}

// New
