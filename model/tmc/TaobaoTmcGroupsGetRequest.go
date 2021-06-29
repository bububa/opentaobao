package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取自定义用户分组列表 API请求
taobao.tmc.groups.get

获取自定义用户分组列表
*/
type TaobaoTmcGroupsGetRequest struct {
    model.Params
    // 要查询分组的名称，多个分组用半角逗号分隔，不传代表查询所有分组信息，但不会返回组下面的用户信息。如果应用没有设置分组则返回空。组名不能以default开头，default开头是系统默认的组。
    groupNames   []string
    // 页码
    pageNo   int64
    // 每页返回多少个分组
    pageSize   int64
}

// 初始化TaobaoTmcGroupsGetRequest对象
func NewTaobaoTmcGroupsGetRequest() *TaobaoTmcGroupsGetRequest{
    return &TaobaoTmcGroupsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTmcGroupsGetRequest) GetApiMethodName() string {
    return "taobao.tmc.groups.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTmcGroupsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GroupNames Setter
// 要查询分组的名称，多个分组用半角逗号分隔，不传代表查询所有分组信息，但不会返回组下面的用户信息。如果应用没有设置分组则返回空。组名不能以default开头，default开头是系统默认的组。
func (r *TaobaoTmcGroupsGetRequest) SetGroupNames(groupNames []string) error {
    r.groupNames = groupNames
    r.Set("group_names", groupNames)
    return nil
}

// GroupNames Getter
func (r TaobaoTmcGroupsGetRequest) GetGroupNames() []string {
    return r.groupNames
}
// PageNo Setter
// 页码
func (r *TaobaoTmcGroupsGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoTmcGroupsGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页返回多少个分组
func (r *TaobaoTmcGroupsGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTmcGroupsGetRequest) GetPageSize() int64 {
    return r.pageSize
}
