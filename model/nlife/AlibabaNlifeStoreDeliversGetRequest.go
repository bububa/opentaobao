package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取门店采购单下的发货单列表 APIRequest
alibaba.nlife.store.delivers.get

获取门店采购单下的发货单列表
*/
type AlibabaNlifeStoreDeliversGetRequest struct {
    model.Params

    // 门店采购订单号
    tradeNo   string 

    // 零售商的门店id
    storeId   int64 

    // 每页的数量
    pageSize   int64 

    // 查询的页码
    pageNo   int64 

}

func NewAlibabaNlifeStoreDeliversGetRequest() *AlibabaNlifeStoreDeliversGetRequest{
    return &AlibabaNlifeStoreDeliversGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaNlifeStoreDeliversGetRequest) GetApiMethodName() string {
    return "alibaba.nlife.store.delivers.get"
}

func (r AlibabaNlifeStoreDeliversGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaNlifeStoreDeliversGetRequest) SetTradeNo(tradeNo string) error {
    r.tradeNo = tradeNo
    r.Set("trade_no", tradeNo)
    return nil
}

func (r AlibabaNlifeStoreDeliversGetRequest) GetTradeNo() string {
    return r.tradeNo
}

func (r *AlibabaNlifeStoreDeliversGetRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaNlifeStoreDeliversGetRequest) GetStoreId() int64 {
    return r.storeId
}

func (r *AlibabaNlifeStoreDeliversGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaNlifeStoreDeliversGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaNlifeStoreDeliversGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r AlibabaNlifeStoreDeliversGetRequest) GetPageNo() int64 {
    return r.pageNo
}

