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
    _context   *WorkBenchContext
    // /workbench/space/application
    _menuUrl   string
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
func (r *AlibabaCampusAclNewListuserbymenuRequest) SetContext(_context *WorkBenchContext) error {
    r._context = _context
    r.Set("context", _context)
    return nil
}

// Context Getter
func (r AlibabaCampusAclNewListuserbymenuRequest) GetContext() *WorkBenchContext {
    return r._context
}
// MenuUrl Setter
// /workbench/space/application
func (r *AlibabaCampusAclNewListuserbymenuRequest) SetMenuUrl(_menuUrl string) error {
    r._menuUrl = _menuUrl
    r.Set("menu_url", _menuUrl)
    return nil
}

// MenuUrl Getter
func (r AlibabaCampusAclNewListuserbymenuRequest) GetMenuUrl() string {
    return r._menuUrl
}
