package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户获取视频互动列表 APIRequest
taobao.interactive.list.getbyuser

根据用户来获取用户编辑的互动列表
*/
type TaobaoInteractiveListGetbyuserRequest struct {
    model.Params

    // 当前页
    currentPage   int64 

    // 每页多少
    pageSize   int64 

}

func NewTaobaoInteractiveListGetbyuserRequest() *TaobaoInteractiveListGetbyuserRequest{
    return &TaobaoInteractiveListGetbyuserRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoInteractiveListGetbyuserRequest) GetApiMethodName() string {
    return "taobao.interactive.list.getbyuser"
}

func (r TaobaoInteractiveListGetbyuserRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoInteractiveListGetbyuserRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r TaobaoInteractiveListGetbyuserRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *TaobaoInteractiveListGetbyuserRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoInteractiveListGetbyuserRequest) GetPageSize() int64 {
    return r.pageSize
}

