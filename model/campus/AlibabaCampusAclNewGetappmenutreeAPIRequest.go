package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclNewGetappmenutreeAPIRequest
查询应用下的菜单树 API请求
alibaba.campus.acl.new.getappmenutree

查询应用下的菜单树 */
type AlibabaCampusAclNewGetappmenutreeAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 是否关联查询出菜单下的权限
	_withpermission bool
}

// New
