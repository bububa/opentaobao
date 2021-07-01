package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclQueryallroleAPIRequest
查询全部角色 API请求
alibaba.campus.acl.queryallrole

查询全部园区 */
type AlibabaCampusAclQueryallroleAPIRequest struct {
	model.Params
	// 公司id不填统一SYS_000
	_companyId int64
	// 系统id
	_systemId string
	// 园区id
	_campusId int64
	// 角色名称
	_roleName string
	// 角色类型
	_roleType string
	// 角色id
	_roleId string
}

// New
