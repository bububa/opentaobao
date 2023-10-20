package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusaclnewsaverolewithmenuAPIRequest 保存角色级联保存角色和权限的关系 API请求
// alibaba.campus.acl.new.saverolewithmenu
//
// 保存角色级联保存角色和权限的关系
type AlibabacampusaclnewsaverolewithmenuAPIRequest struct {
	model.Params
	// 菜单id,权限id
	_treeNodeIds []string
	// 系统入参
	_workbenchcontext *WorkBenchContext
	// 系统自动生成
	_sysRoleDTO *SysRoleDto
}

// NewAlibabacampusaclnewsaverolewithmenuRequest 初始化AlibabacampusaclnewsaverolewithmenuAPIRequest对象
func NewAlibabacampusaclnewsaverolewithmenuRequest() *AlibabacampusaclnewsaverolewithmenuAPIRequest {
	return &AlibabacampusaclnewsaverolewithmenuAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampusaclnewsaverolewithmenuAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.acl.new.saverolewithmenu"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampusaclnewsaverolewithmenuAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampusaclnewsaverolewithmenuAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTreeNodeIds is TreeNodeIds Setter
// 菜单id,权限id
func (r *AlibabacampusaclnewsaverolewithmenuAPIRequest) SetTreeNodeIds(_treeNodeIds []string) error {
	r._treeNodeIds = _treeNodeIds
	r.Set("tree_node_ids", _treeNodeIds)
	return nil
}

// GetTreeNodeIds TreeNodeIds Getter
func (r AlibabacampusaclnewsaverolewithmenuAPIRequest) GetTreeNodeIds() []string {
	return r._treeNodeIds
}

// SetWorkbenchcontext is Workbenchcontext Setter
// 系统入参
func (r *AlibabacampusaclnewsaverolewithmenuAPIRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
	r._workbenchcontext = _workbenchcontext
	r.Set("workbenchcontext", _workbenchcontext)
	return nil
}

// GetWorkbenchcontext Workbenchcontext Getter
func (r AlibabacampusaclnewsaverolewithmenuAPIRequest) GetWorkbenchcontext() *WorkBenchContext {
	return r._workbenchcontext
}

// SetSysRoleDTO is SysRoleDTO Setter
// 系统自动生成
func (r *AlibabacampusaclnewsaverolewithmenuAPIRequest) SetSysRoleDTO(_sysRoleDTO *SysRoleDto) error {
	r._sysRoleDTO = _sysRoleDTO
	r.Set("sys_role_d_t_o", _sysRoleDTO)
	return nil
}

// GetSysRoleDTO SysRoleDTO Getter
func (r AlibabacampusaclnewsaverolewithmenuAPIRequest) GetSysRoleDTO() *SysRoleDto {
	return r._sysRoleDTO
}
