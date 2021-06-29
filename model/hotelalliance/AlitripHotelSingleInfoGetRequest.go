package hotelalliance

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单体酒店信息 API请求
alitrip.hotel.single.info.get

用于给到未来酒店读取与飞猪酒店合作的单体酒店信息，开展单体联盟业务
*/
type AlitripHotelSingleInfoGetRequest struct {
    model.Params
    // 查询酒店信息query参数
    _queryHotelInfoParam   *QueryHotelInfoParam
}

// 初始化AlitripHotelSingleInfoGetRequest对象
func NewAlitripHotelSingleInfoGetRequest() *AlitripHotelSingleInfoGetRequest{
    return &AlitripHotelSingleInfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelSingleInfoGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.single.info.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelSingleInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryHotelInfoParam Setter
// 查询酒店信息query参数
func (r *AlitripHotelSingleInfoGetRequest) SetQueryHotelInfoParam(_queryHotelInfoParam *QueryHotelInfoParam) error {
    r._queryHotelInfoParam = _queryHotelInfoParam
    r.Set("query_hotel_info_param", _queryHotelInfoParam)
    return nil
}

// QueryHotelInfoParam Getter
func (r AlitripHotelSingleInfoGetRequest) GetQueryHotelInfoParam() *QueryHotelInfoParam {
    return r._queryHotelInfoParam
}
