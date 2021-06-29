package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家的分组 API请求
taobao.crm.groups.get

查询卖家的分组，返回查询到的分组列表，分页返回分组
*/
type TaobaoCrmGroupsGetRequest struct {
    model.Params
    // 每页显示的记录数，其最大值不能超过100条，最小值为1，默认20条
    _pageSize   int64
    // 显示第几页的分组，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页码为1
    _currentPage   int64
}

// 初始化TaobaoCrmGroupsGetRequest对象
func NewTaobaoCrmGroupsGetRequest() *TaobaoCrmGroupsGetRequest{
    return &TaobaoCrmGroupsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoCrmGroupsGetRequest) GetApiMethodName() string {
    return "taobao.crm.groups.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoCrmGroupsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageSize Setter
// 每页显示的记录数，其最大值不能超过100条，最小值为1，默认20条
func (r *TaobaoCrmGroupsGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoCrmGroupsGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// CurrentPage Setter
// 显示第几页的分组，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页码为1
func (r *TaobaoCrmGroupsGetRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoCrmGroupsGetRequest) GetCurrentPage() int64 {
    return r._currentPage
}
