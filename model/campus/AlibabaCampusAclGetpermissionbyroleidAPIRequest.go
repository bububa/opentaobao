package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclGetpermissionbyroleidAPIRequest
根据角色Id查询权限 API请求
alibaba.campus.acl.getpermissionbyroleid

根据角色查询权限 */
type AlibabaCampusAclGetpermissionbyroleidAPIRequest struct {
	model.Params
	// 系统id
	_systemId string
	// 园区id
	_campusId int64
	// 角色id
	_roleId string
	// 公司id
	_companyId int64
}

// New
