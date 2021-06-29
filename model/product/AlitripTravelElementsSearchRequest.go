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
    _sellerId   int64
    // 查询关键词
    _query   string
    // 查询数量，限制100
    _count   int64
    // 资源类型
    _type   int64
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
func (r *AlitripTravelElementsSearchRequest) SetSellerId(_sellerId int64) error {
    r._sellerId = _sellerId
    r.Set("seller_id", _sellerId)
    return nil
}

// SellerId Getter
func (r AlitripTravelElementsSearchRequest) GetSellerId() int64 {
    return r._sellerId
}
// Query Setter
// 查询关键词
func (r *AlitripTravelElementsSearchRequest) SetQuery(_query string) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r AlitripTravelElementsSearchRequest) GetQuery() string {
    return r._query
}
// Count Setter
// 查询数量，限制100
func (r *AlitripTravelElementsSearchRequest) SetCount(_count int64) error {
    r._count = _count
    r.Set("count", _count)
    return nil
}

// Count Getter
func (r AlitripTravelElementsSearchRequest) GetCount() int64 {
    return r._count
}
// Type Setter
// 资源类型
func (r *AlitripTravelElementsSearchRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlitripTravelElementsSearchRequest) GetType() int64 {
    return r._type
}
