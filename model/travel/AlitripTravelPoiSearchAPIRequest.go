package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
POI信息查询 API请求
alitrip.travel.poi.search

POI信息查询，用于商品更新使用
*/
type AlitripTravelPoiSearchAPIRequest struct {
    model.Params
    // POI信息ID；
    _id   int64
    // POI信息名称
    _name   string
    // POI类型；1->机场,2->火车站,3->汽车站,4->酒店,5->景点,6->购物，7->美食，9->玩乐，10->阿里旅行服务站，11->服务，100->其他
    _type   int64
}

// 初始化AlitripTravelPoiSearchAPIRequest对象
func NewAlitripTravelPoiSearchRequest() *AlitripTravelPoiSearchAPIRequest{
    return &AlitripTravelPoiSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelPoiSearchAPIRequest) GetApiMethodName() string {
    return "alitrip.travel.poi.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelPoiSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// POI信息ID；
func (r *AlitripTravelPoiSearchAPIRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AlitripTravelPoiSearchAPIRequest) GetId() int64 {
    return r._id
}
// Name Setter
// POI信息名称
func (r *AlitripTravelPoiSearchAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r AlitripTravelPoiSearchAPIRequest) GetName() string {
    return r._name
}
// Type Setter
// POI类型；1->机场,2->火车站,3->汽车站,4->酒店,5->景点,6->购物，7->美食，9->玩乐，10->阿里旅行服务站，11->服务，100->其他
func (r *AlitripTravelPoiSearchAPIRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlitripTravelPoiSearchAPIRequest) GetType() int64 {
    return r._type
}
