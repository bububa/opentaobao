package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclNewListusermenuAPIRequest
查询用户有权限的菜单树 API请求
alibaba.campus.acl.new.listusermenu

查询用户有权限的菜单树 */
type AlibabaCampusAclNewListusermenuAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 用户账号
	_userId string
}

// New
