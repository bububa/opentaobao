package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
组团活动信息查询 APIRequest
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

func NewTaobaoOpentradeGroupActivtiyQueryRequest() *TaobaoOpentradeGroupActivtiyQueryRequest{
    return &TaobaoOpentradeGroupActivtiyQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpentradeGroupActivtiyQueryRequest) GetApiMethodName() string {
    return "taobao.opentrade.group.activtiy.query"
}

func (r TaobaoOpentradeGroupActivtiyQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpentradeGroupActivtiyQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoOpentradeGroupActivtiyQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoOpentradeGroupActivtiyQueryRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r TaobaoOpentradeGroupActivtiyQueryRequest) GetItemId() int64 {
    return r.itemId
}

func (r *TaobaoOpentradeGroupActivtiyQueryRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

func (r TaobaoOpentradeGroupActivtiyQueryRequest) GetPageIndex() int64 {
    return r.pageIndex
}

func (r *TaobaoOpentradeGroupActivtiyQueryRequest) SetGroupActivityId(groupActivityId int64) error {
    r.groupActivityId = groupActivityId
    r.Set("group_activity_id", groupActivityId)
    return nil
}

func (r TaobaoOpentradeGroupActivtiyQueryRequest) GetGroupActivityId() int64 {
    return r.groupActivityId
}

