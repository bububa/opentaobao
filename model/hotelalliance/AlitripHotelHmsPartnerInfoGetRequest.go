package hotelalliance

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取合作商信息 API请求
alitrip.hotel.hms.partner.info.get

用于给到未来酒店读取与飞猪酒店合作的合作商信息，开展单体联盟业务
*/
type AlitripHotelHmsPartnerInfoGetRequest struct {
    model.Params
    // 查询合作商信息query参数
    queryPartnerInfoParam   *QueryPartnerInfoParam
}

// 初始化AlitripHotelHmsPartnerInfoGetRequest对象
func NewAlitripHotelHmsPartnerInfoGetRequest() *AlitripHotelHmsPartnerInfoGetRequest{
    return &AlitripHotelHmsPartnerInfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelHmsPartnerInfoGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.hms.partner.info.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelHmsPartnerInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QueryPartnerInfoParam Setter
// 查询合作商信息query参数
func (r *AlitripHotelHmsPartnerInfoGetRequest) SetQueryPartnerInfoParam(queryPartnerInfoParam *QueryPartnerInfoParam) error {
    r.queryPartnerInfoParam = queryPartnerInfoParam
    r.Set("query_partner_info_param", queryPartnerInfoParam)
    return nil
}

// QueryPartnerInfoParam Getter
func (r AlitripHotelHmsPartnerInfoGetRequest) GetQueryPartnerInfoParam() *QueryPartnerInfoParam {
    return r.queryPartnerInfoParam
}
