package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclNewCheckusermenuAPIRequest
校验用户是否有菜单权限 API请求
alibaba.campus.acl.new.checkusermenu

校验用户是否有菜单权限 */
type AlibabaCampusAclNewCheckusermenuAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 入参
	_param *CheckUserMenuParam
}

// New
