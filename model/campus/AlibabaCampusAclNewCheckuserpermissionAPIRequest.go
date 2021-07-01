package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclNewCheckuserpermissionAPIRequest
校验用户是否有权限 API请求
alibaba.campus.acl.new.checkuserpermission

校验用户是否有权限 */
type AlibabaCampusAclNewCheckuserpermissionAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 接口入参
	_checkUserPermissionParam *CheckUserPermissionParam
}

// New
