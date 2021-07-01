package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclNewListuserrolesAPIRequest
查询并标记用户选择的角色 API请求
alibaba.campus.acl.new.listuserroles

查询并标记用户选择的角色 */
type AlibabaCampusAclNewListuserrolesAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 入参
	_param *UserRoleQueryParam
}

// New
