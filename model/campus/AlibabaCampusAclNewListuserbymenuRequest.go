package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询菜单下的人员 API请求
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

// 初始化AlibabaCampusAclNewListuserbymenuRequest对象
func NewAlibabaCampusAclNewListuserbymenuRequest() *AlibabaCampusAclNewListuserbymenuRequest{
    return &AlibabaCampusAclNewListuserbymenuRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewListuserbymenuRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.listuserbymenu"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewListuserbymenuRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Context Setter
// 系统入参
func (r *AlibabaCampusAclNewListuserbymenuRequest) SetContext(context *WorkBenchContext) error {
    r.context = context
    r.Set("context", context)
    return nil
}

// Context Getter
func (r AlibabaCampusAclNewListuserbymenuRequest) GetContext() *WorkBenchContext {
    return r.context
}
// MenuUrl Setter
// /workbench/space/application
func (r *AlibabaCampusAclNewListuserbymenuRequest) SetMenuUrl(menuUrl string) error {
    r.menuUrl = menuUrl
    r.Set("menu_url", menuUrl)
    return nil
}

// MenuUrl Getter
func (r AlibabaCampusAclNewListuserbymenuRequest) GetMenuUrl() string {
    return r.menuUrl
}
