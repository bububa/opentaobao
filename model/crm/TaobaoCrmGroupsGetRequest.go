package crm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家的分组 APIRequest
taobao.crm.groups.get

查询卖家的分组，返回查询到的分组列表，分页返回分组
*/
type TaobaoCrmGroupsGetRequest struct {
    model.Params

    // 每页显示的记录数，其最大值不能超过100条，最小值为1，默认20条
    pageSize   int64 

    // 显示第几页的分组，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页码为1
    currentPage   int64 

}

func NewTaobaoCrmGroupsGetRequest() *TaobaoCrmGroupsGetRequest{
    return &TaobaoCrmGroupsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCrmGroupsGetRequest) GetApiMethodName() string {
    return "taobao.crm.groups.get"
}

func (r TaobaoCrmGroupsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoCrmGroupsGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoCrmGroupsGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoCrmGroupsGetRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r TaobaoCrmGroupsGetRequest) GetCurrentPage() int64 {
    return r.currentPage
}

