package perfect

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品完美履约信息查询 APIRequest
alibaba.perfect.performance.item.query

同城零售商品完美履约信息查询
*/
type AlibabaPerfectPerformanceItemQueryRequest struct {
    model.Params

    // 查询入参
    itemPerfectPerformanceQueryReq   *ItemPerfectPerformanceQueryReq 

}

func NewAlibabaPerfectPerformanceItemQueryRequest() *AlibabaPerfectPerformanceItemQueryRequest{
    return &AlibabaPerfectPerformanceItemQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaPerfectPerformanceItemQueryRequest) GetApiMethodName() string {
    return "alibaba.perfect.performance.item.query"
}

func (r AlibabaPerfectPerformanceItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaPerfectPerformanceItemQueryRequest) SetItemPerfectPerformanceQueryReq(itemPerfectPerformanceQueryReq *ItemPerfectPerformanceQueryReq) error {
    r.itemPerfectPerformanceQueryReq = itemPerfectPerformanceQueryReq
    r.Set("item_perfect_performance_query_req", itemPerfectPerformanceQueryReq)
    return nil
}

func (r AlibabaPerfectPerformanceItemQueryRequest) GetItemPerfectPerformanceQueryReq() *ItemPerfectPerformanceQueryReq {
    return r.itemPerfectPerformanceQueryReq
}

