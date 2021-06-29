package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户有权限的菜单树 APIRequest
alibaba.campus.acl.new.listusermenu

查询用户有权限的菜单树
*/
type AlibabaCampusAclNewListusermenuRequest struct {
    model.Params

    // 系统入参
    workbenchcontext   *WorkBenchContext 

    // 用户账号
    userId   string 

}

func NewAlibabaCampusAclNewListusermenuRequest() *AlibabaCampusAclNewListusermenuRequest{
    return &AlibabaCampusAclNewListusermenuRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclNewListusermenuRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.listusermenu"
}

func (r AlibabaCampusAclNewListusermenuRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclNewListusermenuRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

func (r AlibabaCampusAclNewListusermenuRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}

func (r *AlibabaCampusAclNewListusermenuRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaCampusAclNewListusermenuRequest) GetUserId() string {
    return r.userId
}

