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
type AlibabaPerfectPerformanceItemQueryAPIRequest struct {
    model.Params
    // 查询入参
    _itemPerfectPerformanceQueryReq   *ItemPerfectPerformanceQueryReq
}

// 初始化AlibabaPerfectPerformanceItemQueryAPIRequest对象
func NewAlibabaPerfectPerformanceItemQueryRequest() *AlibabaPerfectPerformanceItemQueryAPIRequest{
    return &AlibabaPerfectPerformanceItemQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPerfectPerformanceItemQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.perfect.performance.item.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPerfectPerformanceItemQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemPerfectPerformanceQueryReq Setter
// 查询入参
func (r *AlibabaPerfectPerformanceItemQueryAPIRequest) SetItemPerfectPerformanceQueryReq(_itemPerfectPerformanceQueryReq *ItemPerfectPerformanceQueryReq) error {
    r._itemPerfectPerformanceQueryReq = _itemPerfectPerformanceQueryReq
    r.Set("item_perfect_performance_query_req", _itemPerfectPerformanceQueryReq)
    return nil
}

// ItemPerfectPerformanceQueryReq Getter
func (r AlibabaPerfectPerformanceItemQueryAPIRequest) GetItemPerfectPerformanceQueryReq() *ItemPerfectPerformanceQueryReq {
    return r._itemPerfectPerformanceQueryReq
}
