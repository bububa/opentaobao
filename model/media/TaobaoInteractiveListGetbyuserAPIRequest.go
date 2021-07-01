package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户获取视频互动列表 API请求
taobao.interactive.list.getbyuser

根据用户来获取用户编辑的互动列表
*/
type TaobaoInteractiveListGetbyuserAPIRequest struct {
    model.Params
    // 当前页
    _currentPage   int64
    // 每页多少
    _pageSize   int64
}

// 初始化TaobaoInteractiveListGetbyuserAPIRequest对象
func NewTaobaoInteractiveListGetbyuserRequest() *TaobaoInteractiveListGetbyuserAPIRequest{
    return &TaobaoInteractiveListGetbyuserAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoInteractiveListGetbyuserAPIRequest) GetApiMethodName() string {
    return "taobao.interactive.list.getbyuser"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoInteractiveListGetbyuserAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CurrentPage Setter
// 当前页
func (r *TaobaoInteractiveListGetbyuserAPIRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r TaobaoInteractiveListGetbyuserAPIRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// PageSize Setter
// 每页多少
func (r *TaobaoInteractiveListGetbyuserAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoInteractiveListGetbyuserAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
