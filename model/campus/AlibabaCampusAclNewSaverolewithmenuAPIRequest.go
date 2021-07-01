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

// NewAlibabaCampusAclNewSaverolewithmenuRequest 初始化AlibabaCampusAclNewSaverolewithmenuAPIRequest对象
func NewAlibabaCampusAclNewSaverolewithmenuRequest() *AlibabaCampusAclNewSaverolewithmenuAPIRequest {
	return &AlibabaCampusAclNewSaverolewithmenuAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewSaverolewithmenuAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.saverolewithmenu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewSaverolewithmenuAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewSaverolewithmenuAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// Get Workbenchcontext Getter
func (r AlibabaCampusAclNewSaverolewithmenuAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// Set is SysRoleDTO Setter
// 系统自动生成
func (r *AlibabaCampusAclNewSaverolewithmenuAPIRequest) SetSysRoleDTO(_sysRoleDTO *SysRoleDto) error {
	r._sysRoleDTO = _sysRoleDTO
	r.Set("sys_role_d_t_o", _sysRoleDTO)
	return nil
}

// Get SysRoleDTO Getter
func (r AlibabaCampusAclNewSaverolewithmenuAPIRequest) GetSysRoleDTO() *SysRoleDto {
	return r._sysRoleDTO
}

// Set is TreeNodeIds Setter
// 菜单id,权限id
func (r *AlibabaCampusAclNewSaverolewithmenuAPIRequest) SetTreeNodeIds(_treeNodeIds []string) error {
	r._treeNodeIds = _treeNodeIds
	r.Set("tree_node_ids", _treeNodeIds)
	return nil
}

// Get TreeNodeIds Getter
func (r AlibabaCampusAclNewSaverolewithmenuAPIRequest) GetTreeNodeIds() []string {
	return r._treeNodeIds
}
