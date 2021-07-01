package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusAclNewSaverolewithmenuAPIRequest
保存角色级联保存角色和权限的关系 API请求
alibaba.campus.acl.new.saverolewithmenu

保存角色级联保存角色和权限的关系 */
type AlibabaCampusAclNewSaverolewithmenuAPIRequest struct {
	model.Params
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 系统自动生成
	_sysRoleDTO *SysRoleDto
	// 菜单id,权限id
	_treeNodeIds []string
}

// New
