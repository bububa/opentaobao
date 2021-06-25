package hotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店搜索List接口 APIRequest
alitrip.hotel.search.list.get

酒店搜索List接口
*/
type AlitripHotelSearchListGetRequest struct {
    model.Params

    // 入参
    paramTopHotelSearchListParam   *TopHotelSearchListParam 

}

func NewAlitripHotelSearchListGetRequest() *AlitripHotelSearchListGetRequest{
    return &AlitripHotelSearchListGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripHotelSearchListGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.search.list.get"
}

func (r AlitripHotelSearchListGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripHotelSearchListGetRequest) SetParamTopHotelSearchListParam(paramTopHotelSearchListParam *TopHotelSearchListParam) error {
    r.paramTopHotelSearchListParam = paramTopHotelSearchListParam
    r.Set("param_top_hotel_search_list_param", paramTopHotelSearchListParam)
    return nil
}

func (r AlitripHotelSearchListGetRequest) GetParamTopHotelSearchListParam() *TopHotelSearchListParam {
    return r.paramTopHotelSearchListParam
}

