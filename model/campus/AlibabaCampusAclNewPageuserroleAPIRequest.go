package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclNewPageuserroleAPIRequest
分页查询管理员 API请求
alibaba.campus.acl.new.pageuserrole

新增用户和角色的关系 */
type AlibabaCampusAclNewPageuserroleAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 入参
	_usersRoleQueryParam *UsersRoleQueryParam
}

// New
