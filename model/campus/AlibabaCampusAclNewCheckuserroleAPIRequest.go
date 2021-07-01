package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclNewCheckuserroleAPIRequest
校验用户是否有角色 API请求
alibaba.campus.acl.new.checkuserrole

校验用户是否有角色 */
type AlibabaCampusAclNewCheckuserroleAPIRequest struct {
	model.Params
	// 用户账号
	_userId string
	// 角色key
	_roleKey string
	// 系统入参
	_workbenchcontext *WorkBenchContext
}

// New
