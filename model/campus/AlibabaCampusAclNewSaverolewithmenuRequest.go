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
    workbenchcontext   *WorkBenchContext
    // 系统自动生成
    sysRoleDTO   *SysRoleDto
    // 菜单id,权限id
    treeNodeIds   []string
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
func (r *AlibabaCampusAclNewSaverolewithmenuRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

// Workbenchcontext Getter
func (r AlibabaCampusAclNewSaverolewithmenuRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}
// SysRoleDTO Setter
// 系统自动生成
func (r *AlibabaCampusAclNewSaverolewithmenuRequest) SetSysRoleDTO(sysRoleDTO *SysRoleDto) error {
    r.sysRoleDTO = sysRoleDTO
    r.Set("sys_role_d_t_o", sysRoleDTO)
    return nil
}

// SysRoleDTO Getter
func (r AlibabaCampusAclNewSaverolewithmenuRequest) GetSysRoleDTO() *SysRoleDto {
    return r.sysRoleDTO
}
// TreeNodeIds Setter
// 菜单id,权限id
func (r *AlibabaCampusAclNewSaverolewithmenuRequest) SetTreeNodeIds(treeNodeIds []string) error {
    r.treeNodeIds = treeNodeIds
    r.Set("tree_node_ids", treeNodeIds)
    return nil
}

// TreeNodeIds Getter
func (r AlibabaCampusAclNewSaverolewithmenuRequest) GetTreeNodeIds() []string {
    return r.treeNodeIds
}
