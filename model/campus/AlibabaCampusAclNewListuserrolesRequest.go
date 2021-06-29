package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询并标记用户选择的角色 API请求
alibaba.campus.acl.new.listuserroles

查询并标记用户选择的角色
*/
type AlibabaCampusAclNewListuserrolesRequest struct {
    model.Params
    // 系统入参
    workbenchcontext   *WorkBenchContext
    // 入参
    param   *UserRoleQueryParam
}

// 初始化AlibabaCampusAclNewListuserrolesRequest对象
func NewAlibabaCampusAclNewListuserrolesRequest() *AlibabaCampusAclNewListuserrolesRequest{
    return &AlibabaCampusAclNewListuserrolesRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewListuserrolesRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.listuserroles"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewListuserrolesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewListuserrolesRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

// Workbenchcontext Getter
func (r AlibabaCampusAclNewListuserrolesRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}
// Param Setter
// 入参
func (r *AlibabaCampusAclNewListuserrolesRequest) SetParam(param *UserRoleQueryParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaCampusAclNewListuserrolesRequest) GetParam() *UserRoleQueryParam {
    return r.param
}
