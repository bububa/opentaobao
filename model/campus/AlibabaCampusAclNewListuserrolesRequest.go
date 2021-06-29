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
    _workbenchcontext   *WorkBenchContext
    // 入参
    _param   *UserRoleQueryParam
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
func (r *AlibabaCampusAclNewListuserrolesRequest) SetWorkbenchcontext(_workbenchcontext *WorkBenchContext) error {
    r._workbenchcontext = _workbenchcontext
    r.Set("workbenchcontext", _workbenchcontext)
    return nil
}

// Workbenchcontext Getter
func (r AlibabaCampusAclNewListuserrolesRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r._workbenchcontext
}
// Param Setter
// 入参
func (r *AlibabaCampusAclNewListuserrolesRequest) SetParam(_param *UserRoleQueryParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaCampusAclNewListuserrolesRequest) GetParam() *UserRoleQueryParam {
    return r._param
}