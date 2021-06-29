package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家元素搜索 API请求
alitrip.travel.elements.search

提供商家维护的景点、酒店、餐饮等元素搜索
*/
type AlitripTravelElementsSearchRequest struct {
    model.Params
    // 商家id
    sellerId   int64
    // 查询关键词
    query   string
    // 查询数量，限制100
    count   int64
    // 资源类型
    type   int64
}

// 初始化AlitripTravelElementsSearchRequest对象
func NewAlitripTravelElementsSearchRequest() *AlitripTravelElementsSearchRequest{
    return &AlitripTravelElementsSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelElementsSearchRequest) GetApiMethodName() string {
    return "alitrip.travel.elements.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelElementsSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SellerId Setter
// 商家id
func (r *AlitripTravelElementsSearchRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

// SellerId Getter
func (r AlitripTravelElementsSearchRequest) GetSellerId() int64 {
    return r.sellerId
}
// Query Setter
// 查询关键词
func (r *AlitripTravelElementsSearchRequest) SetQuery(query string) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r AlitripTravelElementsSearchRequest) GetQuery() string {
    return r.query
}
// Count Setter
// 查询数量，限制100
func (r *AlitripTravelElementsSearchRequest) SetCount(count int64) error {
    r.count = count
    r.Set("count", count)
    return nil
}

// Count Getter
func (r AlitripTravelElementsSearchRequest) GetCount() int64 {
    return r.count
}
// Type Setter
// 资源类型
func (r *AlitripTravelElementsSearchRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlitripTravelElementsSearchRequest) GetType() int64 {
    return r.type
}
