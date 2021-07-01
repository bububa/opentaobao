package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclNewListrolesAPIRequest
查询全部角色 API请求
alibaba.campus.acl.new.listroles

查询全部角色 */
type AlibabaCampusAclNewListrolesAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 入参
	_rolequeryparam *RoleQueryParam
}

// New
