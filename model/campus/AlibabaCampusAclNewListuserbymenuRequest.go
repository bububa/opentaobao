package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询菜单下的人员 APIRequest
alibaba.campus.acl.new.listuserbymenu

查询拥有菜单权限的用户
*/
type AlibabaCampusAclNewListuserbymenuRequest struct {
    model.Params

    // 系统入参
    context   *WorkBenchContext 

    // /workbench/space/application
    menuUrl   string 

}

func NewAlibabaCampusAclNewListuserbymenuRequest() *AlibabaCampusAclNewListuserbymenuRequest{
    return &AlibabaCampusAclNewListuserbymenuRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusAclNewListuserbymenuRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.listuserbymenu"
}

func (r AlibabaCampusAclNewListuserbymenuRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusAclNewListuserbymenuRequest) SetContext(context *WorkBenchContext) error {
    r.context = context
    r.Set("context", context)
    return nil
}

func (r AlibabaCampusAclNewListuserbymenuRequest) GetContext() *WorkBenchContext {
    return r.context
}

func (r *AlibabaCampusAclNewListuserbymenuRequest) SetMenuUrl(menuUrl string) error {
    r.menuUrl = menuUrl
    r.Set("menu_url", menuUrl)
    return nil
}

func (r AlibabaCampusAclNewListuserbymenuRequest) GetMenuUrl() string {
    return r.menuUrl
}

