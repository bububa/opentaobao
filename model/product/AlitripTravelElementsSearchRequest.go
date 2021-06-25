package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家元素搜索 APIRequest
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

func NewAlitripTravelElementsSearchRequest() *AlitripTravelElementsSearchRequest{
    return &AlitripTravelElementsSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTravelElementsSearchRequest) GetApiMethodName() string {
    return "alitrip.travel.elements.search"
}

func (r AlitripTravelElementsSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTravelElementsSearchRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

func (r AlitripTravelElementsSearchRequest) GetSellerId() int64 {
    return r.sellerId
}

func (r *AlitripTravelElementsSearchRequest) SetQuery(query string) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlitripTravelElementsSearchRequest) GetQuery() string {
    return r.query
}

func (r *AlitripTravelElementsSearchRequest) SetCount(count int64) error {
    r.count = count
    r.Set("count", count)
    return nil
}

func (r AlitripTravelElementsSearchRequest) GetCount() int64 {
    return r.count
}

func (r *AlitripTravelElementsSearchRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlitripTravelElementsSearchRequest) GetType() int64 {
    return r.type
}

