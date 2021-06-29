package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取奇门用户列表 API请求
taobao.qimen.trade.users.get

获取已开通奇门订单服务的用户列表
*/
type TaobaoQimenTradeUsersGetRequest struct {
    model.Params
    // 每页的数量
    _pageIndex   int64
    // 页数
    _pageSize   int64
}

// 初始化TaobaoQimenTradeUsersGetRequest对象
func NewTaobaoQimenTradeUsersGetRequest() *TaobaoQimenTradeUsersGetRequest{
    return &TaobaoQimenTradeUsersGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenTradeUsersGetRequest) GetApiMethodName() string {
    return "taobao.qimen.trade.users.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenTradeUsersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageIndex Setter
// 每页的数量
func (r *TaobaoQimenTradeUsersGetRequest) SetPageIndex(_pageIndex int64) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoQimenTradeUsersGetRequest) GetPageIndex() int64 {
    return r._pageIndex
}
// PageSize Setter
// 页数
func (r *TaobaoQimenTradeUsersGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoQimenTradeUsersGetRequest) GetPageSize() int64 {
    return r._pageSize
}
