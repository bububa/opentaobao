package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组团活动信息查询 API请求
taobao.opentrade.group.activtiy.query

组团购场景下，团购活动信息查询
*/
type TaobaoOpentradeGroupActivtiyQueryRequest struct {
    model.Params
    // 分页参数，每页大小
    _pageSize   int64
    // 商品ID
    _itemId   int64
    // 分页参数，当前页，以0开始
    _pageIndex   int64
    // 组团活动id
    _groupActivityId   int64
}

// 初始化TaobaoOpentradeGroupActivtiyQueryRequest对象
func NewTaobaoOpentradeGroupActivtiyQueryRequest() *TaobaoOpentradeGroupActivtiyQueryRequest{
    return &TaobaoOpentradeGroupActivtiyQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeGroupActivtiyQueryRequest) GetApiMethodName() string {
    return "taobao.opentrade.group.activtiy.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeGroupActivtiyQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageSize Setter
// 分页参数，每页大小
func (r *TaobaoOpentradeGroupActivtiyQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOpentradeGroupActivtiyQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
// ItemId Setter
// 商品ID
func (r *TaobaoOpentradeGroupActivtiyQueryRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOpentradeGroupActivtiyQueryRequest) GetItemId() int64 {
    return r._itemId
}
// PageIndex Setter
// 分页参数，当前页，以0开始
func (r *TaobaoOpentradeGroupActivtiyQueryRequest) SetPageIndex(_pageIndex int64) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoOpentradeGroupActivtiyQueryRequest) GetPageIndex() int64 {
    return r._pageIndex
}
// GroupActivityId Setter
// 组团活动id
func (r *TaobaoOpentradeGroupActivtiyQueryRequest) SetGroupActivityId(_groupActivityId int64) error {
    r._groupActivityId = _groupActivityId
    r.Set("group_activity_id", _groupActivityId)
    return nil
}

// GroupActivityId Getter
func (r TaobaoOpentradeGroupActivtiyQueryRequest) GetGroupActivityId() int64 {
    return r._groupActivityId
}
