package hotelalliance

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取合作商信息 APIRequest
alitrip.hotel.hms.partner.info.get

用于给到未来酒店读取与飞猪酒店合作的合作商信息，开展单体联盟业务
*/
type AlitripHotelHmsPartnerInfoGetRequest struct {
    model.Params

    // 查询合作商信息query参数
    queryPartnerInfoParam   *QueryPartnerInfoParam 

}

func NewAlitripHotelHmsPartnerInfoGetRequest() *AlitripHotelHmsPartnerInfoGetRequest{
    return &AlitripHotelHmsPartnerInfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripHotelHmsPartnerInfoGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.hms.partner.info.get"
}

func (r AlitripHotelHmsPartnerInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripHotelHmsPartnerInfoGetRequest) SetQueryPartnerInfoParam(queryPartnerInfoParam *QueryPartnerInfoParam) error {
    r.queryPartnerInfoParam = queryPartnerInfoParam
    r.Set("query_partner_info_param", queryPartnerInfoParam)
    return nil
}

func (r AlitripHotelHmsPartnerInfoGetRequest) GetQueryPartnerInfoParam() *QueryPartnerInfoParam {
    return r.queryPartnerInfoParam
}

