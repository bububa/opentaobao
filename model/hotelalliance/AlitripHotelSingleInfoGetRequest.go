package hotelalliance

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单体酒店信息 APIRequest
alitrip.hotel.single.info.get

用于给到未来酒店读取与飞猪酒店合作的单体酒店信息，开展单体联盟业务
*/
type AlitripHotelSingleInfoGetRequest struct {
    model.Params

    // 查询酒店信息query参数
    queryHotelInfoParam   *QueryHotelInfoParam 

}

func NewAlitripHotelSingleInfoGetRequest() *AlitripHotelSingleInfoGetRequest{
    return &AlitripHotelSingleInfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripHotelSingleInfoGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.single.info.get"
}

func (r AlitripHotelSingleInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripHotelSingleInfoGetRequest) SetQueryHotelInfoParam(queryHotelInfoParam *QueryHotelInfoParam) error {
    r.queryHotelInfoParam = queryHotelInfoParam
    r.Set("query_hotel_info_param", queryHotelInfoParam)
    return nil
}

func (r AlitripHotelSingleInfoGetRequest) GetQueryHotelInfoParam() *QueryHotelInfoParam {
    return r.queryHotelInfoParam
}

