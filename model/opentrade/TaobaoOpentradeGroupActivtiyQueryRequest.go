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
    pageSize   int64
    // 商品ID
    itemId   int64
    // 分页参数，当前页，以0开始
    pageIndex   int64
    // 组团活动id
    groupActivityId   int64
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
func (r *TaobaoOpentradeGroupActivtiyQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOpentradeGroupActivtiyQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
// ItemId Setter
// 商品ID
func (r *TaobaoOpentradeGroupActivtiyQueryRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r TaobaoOpentradeGroupActivtiyQueryRequest) GetItemId() int64 {
    return r.itemId
}
// PageIndex Setter
// 分页参数，当前页，以0开始
func (r *TaobaoOpentradeGroupActivtiyQueryRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoOpentradeGroupActivtiyQueryRequest) GetPageIndex() int64 {
    return r.pageIndex
}
// GroupActivityId Setter
// 组团活动id
func (r *TaobaoOpentradeGroupActivtiyQueryRequest) SetGroupActivityId(groupActivityId int64) error {
    r.groupActivityId = groupActivityId
    r.Set("group_activity_id", groupActivityId)
    return nil
}

// GroupActivityId Getter
func (r TaobaoOpentradeGroupActivtiyQueryRequest) GetGroupActivityId() int64 {
    return r.groupActivityId
}
