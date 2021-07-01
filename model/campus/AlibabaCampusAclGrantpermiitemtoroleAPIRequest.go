package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclGrantpermiitemtoroleAPIRequest
权限赋予角色 API请求
alibaba.campus.acl.grantpermiitemtorole

权限赋予角色 */
type AlibabaCampusAclGrantpermiitemtoroleAPIRequest struct {
	model.Params
	// 公司ID,不填统一初始化SYS_000
	_companyId int64
	// 系统id
	_systemId string
	// 园区ID
	_campusId int64
	// 系统自动生成
	_role *RoleReq
	// 系统自动生成
	_priv []PermissionReq
	// 操作人id
	_userId string
}

// New
