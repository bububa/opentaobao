package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest
根据角色id查询权限 API请求
alibaba.campus.acl.new.getrolewithmenutreenodes

根据角色id查询权限 */
type AlibabaCampusAclNewGetrolewithmenutreenodesAPIRequest struct {
	model.Params
	// 角色id
	_roleId int64
	// 是否查询全部类型权限
	_allPermission bool
	// 系统参数
	_workbenchcontext *WorkBenchContext
}

// New
