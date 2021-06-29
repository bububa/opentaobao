package perfect

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品完美履约信息查询 API请求
alibaba.perfect.performance.item.query

同城零售商品完美履约信息查询
*/
type AlibabaPerfectPerformanceItemQueryRequest struct {
    model.Params
    // 查询入参
    _itemPerfectPerformanceQueryReq   *ItemPerfectPerformanceQueryReq
}

// 初始化AlibabaPerfectPerformanceItemQueryRequest对象
func NewAlibabaPerfectPerformanceItemQueryRequest() *AlibabaPerfectPerformanceItemQueryRequest{
    return &AlibabaPerfectPerformanceItemQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPerfectPerformanceItemQueryRequest) GetApiMethodName() string {
    return "alibaba.perfect.performance.item.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPerfectPerformanceItemQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemPerfectPerformanceQueryReq Setter
// 查询入参
func (r *AlibabaPerfectPerformanceItemQueryRequest) SetItemPerfectPerformanceQueryReq(_itemPerfectPerformanceQueryReq *ItemPerfectPerformanceQueryReq) error {
    r._itemPerfectPerformanceQueryReq = _itemPerfectPerformanceQueryReq
    r.Set("item_perfect_performance_query_req", _itemPerfectPerformanceQueryReq)
    return nil
}

// ItemPerfectPerformanceQueryReq Getter
func (r AlibabaPerfectPerformanceItemQueryRequest) GetItemPerfectPerformanceQueryReq() *ItemPerfectPerformanceQueryReq {
    return r._itemPerfectPerformanceQueryReq
}
