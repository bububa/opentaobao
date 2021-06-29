package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
保存角色级联保存角色和权限的关系 APIRequest
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

func NewAlibabaCampusAclNewSaverolewithmenuRequest() *AlibabaCampusAclNewSaverolewithmenuRequest{
    return &AlibabaCampusAclNewSaverolewithmenuRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclNewSaverolewithmenuRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.saverolewithmenu"
}

func (r AlibabaCampusAclNewSaverolewithmenuRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclNewSaverolewithmenuRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

func (r AlibabaCampusAclNewSaverolewithmenuRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}

func (r *AlibabaCampusAclNewSaverolewithmenuRequest) SetSysRoleDTO(sysRoleDTO *SysRoleDto) error {
    r.sysRoleDTO = sysRoleDTO
    r.Set("sys_role_d_t_o", sysRoleDTO)
    return nil
}

func (r AlibabaCampusAclNewSaverolewithmenuRequest) GetSysRoleDTO() *SysRoleDto {
    return r.sysRoleDTO
}

func (r *AlibabaCampusAclNewSaverolewithmenuRequest) SetTreeNodeIds(treeNodeIds []string) error {
    r.treeNodeIds = treeNodeIds
    r.Set("tree_node_ids", treeNodeIds)
    return nil
}

func (r AlibabaCampusAclNewSaverolewithmenuRequest) GetTreeNodeIds() []string {
    return r.treeNodeIds
}

