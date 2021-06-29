package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询用户有权限的菜单树 API请求
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

// 初始化AlibabaCampusAclNewListusermenuRequest对象
func NewAlibabaCampusAclNewListusermenuRequest() *AlibabaCampusAclNewListusermenuRequest{
    return &AlibabaCampusAclNewListusermenuRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusAclNewListusermenuRequest) GetApiMethodName() string {
    return "alibaba.campus.acl.new.listusermenu"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusAclNewListusermenuRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Workbenchcontext Setter
// 系统入参
func (r *AlibabaCampusAclNewListusermenuRequest) SetWorkbenchcontext(workbenchcontext *WorkBenchContext) error {
    r.workbenchcontext = workbenchcontext
    r.Set("workbenchcontext", workbenchcontext)
    return nil
}

// Workbenchcontext Getter
func (r AlibabaCampusAclNewListusermenuRequest) GetWorkbenchcontext() *WorkBenchContext {
    return r.workbenchcontext
}
// UserId Setter
// 用户账号
func (r *AlibabaCampusAclNewListusermenuRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaCampusAclNewListusermenuRequest) GetUserId() string {
    return r.userId
}
