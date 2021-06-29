package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否有菜单权限 APIRequest
alibaba.campus.acl.new.checkusermenu

校验用户是否有菜单权限
*/
type AlibabaCampusAclNewCheckusermenuRequest struct {
    model.Params

    // 系统入参
    workbenchcontext   *WorkBenchContext 

    // 入参
    param   *CheckUserMenuParam 

}

func NewAlibabaCampusAclNewCheckusermenuRequest() *AlibabaCampusAclNewCheckusermenuRequest{
    return &AlibabaCampusAclNewCheckusermenuRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclNewCheckusermenuRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.checkusermenu"
}

func (r AlibabaCampusAclNewCheckusermenuRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclNewCheckusermenuRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

func (r AlibabaCampusAclNewCheckusermenuRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}

func (r *AlibabaCampusAclNewCheckusermenuRequest) SetParam(param *CheckUserMenuParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaCampusAclNewCheckusermenuRequest) GetParam() *CheckUserMenuParam {
    return r.param
}

