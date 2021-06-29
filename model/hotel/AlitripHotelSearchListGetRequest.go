package hotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店搜索List接口 API请求
alitrip.hotel.search.list.get

酒店搜索List接口
*/
type AlitripHotelSearchListGetRequest struct {
    model.Params
    // 入参
    paramTopHotelSearchListParam   *TopHotelSearchListParam
}

// 初始化AlitripHotelSearchListGetRequest对象
func NewAlitripHotelSearchListGetRequest() *AlitripHotelSearchListGetRequest{
    return &AlitripHotelSearchListGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelSearchListGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.search.list.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelSearchListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamTopHotelSearchListParam Setter
// 入参
func (r *AlitripHotelSearchListGetRequest) SetParamTopHotelSearchListParam(paramTopHotelSearchListParam *TopHotelSearchListParam) error {
    r.paramTopHotelSearchListParam = paramTopHotelSearchListParam
    r.Set("param_top_hotel_search_list_param", paramTopHotelSearchListParam)
    return nil
}

// ParamTopHotelSearchListParam Getter
func (r AlitripHotelSearchListGetRequest) GetParamTopHotelSearchListParam() *TopHotelSearchListParam {
    return r.paramTopHotelSearchListParam
}
