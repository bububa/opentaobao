package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取奇门用户列表 APIRequest
taobao.qimen.trade.users.get

获取已开通奇门订单服务的用户列表
*/
type TaobaoQimenTradeUsersGetRequest struct {
    model.Params

    // 每页的数量
    pageIndex   int64 

    // 页数
    pageSize   int64 

}

func NewTaobaoQimenTradeUsersGetRequest() *TaobaoQimenTradeUsersGetRequest{
    return &TaobaoQimenTradeUsersGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenTradeUsersGetRequest) GetApiMethodName() string {
    return "taobao.qimen.trade.users.get"
}

func (r TaobaoQimenTradeUsersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenTradeUsersGetRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

func (r TaobaoQimenTradeUsersGetRequest) GetPageIndex() int64 {
    return r.pageIndex
}

func (r *TaobaoQimenTradeUsersGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoQimenTradeUsersGetRequest) GetPageSize() int64 {
    return r.pageSize
}

