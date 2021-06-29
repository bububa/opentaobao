package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
保存角色级联保存角色和权限的关系 API请求
alibaba.campus.acl.new.saverolewithmenu

保存角色级联保存角色和权限的关系
*/
type AlibabaCampusAclNewSaverolewithmenuRequest struct {
    model.Params
    // 系统入参
    _workbenchcontext   *WorkBenchContext
    // 系统自动生成
    _sysRoleDTO   *SysRoleDTO
    // 菜单id,权限id
    _treeNodeIds   []string
}

// 初始化AlibabaCampusAclNewSaverolewithmenuRequest对象
func NewAlibabaCampusAclNewSaverolewithmenuRequest() *AlibabaCampusAclNewSaverolewithmenuRequest{
    return &AlibabaCampusAclNewSaverolewithmenuRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewSaverolewithmenuRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.saverolewithmenu"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewSaverolewithmenuRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewSaverolewithmenuRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
    r._workbenchcontext = _workbenchcontext
    r.Set("workbenchcontext", _workbenchcontext)
    return nil
}

// Workbenchcontext Getter
func (r AlibabaCampusAclNewSaverolewithmenuRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r._workbenchcontext
}
// SysRoleDTO Setter
// 系统自动生成
func (r *AlibabaCampusAclNewSaverolewithmenuRequest) SetSysRoleDTO(_sysRoleDTO *SysRoleDTO) error {
    r._sysRoleDTO = _sysRoleDTO
    r.Set("sys_role_d_t_o", _sysRoleDTO)
    return nil
}

// SysRoleDTO Getter
func (r AlibabaCampusAclNewSaverolewithmenuRequest) GetSysRoleDTO() *SysRoleDTO {
    return r._sysRoleDTO
}
// TreeNodeIds Setter
// 菜单id,权限id
func (r *AlibabaCampusAclNewSaverolewithmenuRequest) SetTreeNodeIds(_treeNodeIds []string) error {
    r._treeNodeIds = _treeNodeIds
    r.Set("tree_node_ids", _treeNodeIds)
    return nil
}

// TreeNodeIds Getter
func (r AlibabaCampusAclNewSaverolewithmenuRequest) GetTreeNodeIds() []string {
    return r._treeNodeIds
}
