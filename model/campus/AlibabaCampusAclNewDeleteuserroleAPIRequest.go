package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclNewDeleteuserroleAPIRequest
删除管理员 API请求
alibaba.campus.acl.new.deleteuserrole

删除管理员 */
type AlibabaCampusAclNewDeleteuserroleAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 用户账号
	_userId string
	// 角色id
	_roleIds []int64
}

// New
