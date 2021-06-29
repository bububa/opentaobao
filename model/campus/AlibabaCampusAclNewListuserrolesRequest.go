package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询并标记用户选择的角色 APIRequest
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

func NewAlibabaCampusAclNewListuserrolesRequest() *AlibabaCampusAclNewListuserrolesRequest{
    return &AlibabaCampusAclNewListuserrolesRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclNewListuserrolesRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.listuserroles"
}

func (r AlibabaCampusAclNewListuserrolesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclNewListuserrolesRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

func (r AlibabaCampusAclNewListuserrolesRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}

func (r *AlibabaCampusAclNewListuserrolesRequest) SetParam(param *UserRoleQueryParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaCampusAclNewListuserrolesRequest) GetParam() *UserRoleQueryParam {
    return r.param
}

