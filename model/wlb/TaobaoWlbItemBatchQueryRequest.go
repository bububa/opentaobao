package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批次库存查询接口 APIRequest
taobao.wlb.item.batch.query

根据用户id，item id list和store code来查询商品库存信息和批次信息
*/
type TaobaoWlbItemBatchQueryRequest struct {
    model.Params

    // 仓库编号
    storeCode   string 

    // 分页查询参数，指定查询页数，默认为1
    pageNo   int64 

    // 分页查询参数，每页查询数量，默认20，最大值50,大于50时按照50条查询
    pageSize   int64 

    // 需要查询的商品ID列表，以字符串表示，ID间以;隔开
    itemIds   string 

}

func NewTaobaoWlbItemBatchQueryRequest() *TaobaoWlbItemBatchQueryRequest{
    return &TaobaoWlbItemBatchQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbItemBatchQueryRequest) GetApiMethodName() string {
    return "taobao.wlb.item.batch.query"
}

func (r TaobaoWlbItemBatchQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbItemBatchQueryRequest) SetStoreCode(storeCode string) error {
    r.storeCode = storeCode
    r.Set("store_code", storeCode)
    return nil
}

func (r TaobaoWlbItemBatchQueryRequest) GetStoreCode() string {
    return r.storeCode
}

func (r *TaobaoWlbItemBatchQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoWlbItemBatchQueryRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoWlbItemBatchQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoWlbItemBatchQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoWlbItemBatchQueryRequest) SetItemIds(itemIds string) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

func (r TaobaoWlbItemBatchQueryRequest) GetItemIds() string {
    return r.itemIds
}

