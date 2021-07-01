package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclNewListuserbymenuAPIRequest
查询菜单下的人员 API请求
alibaba.campus.acl.new.listuserbymenu

查询拥有菜单权限的用户 */
type AlibabaCampusAclNewListuserbymenuAPIRequest struct {
	model.Params
	// 系统入参
	_context *WorkBenchContext
	// /workbench/space/application
	_menuUrl string
}

// New
