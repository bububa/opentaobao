package hotelalliance

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取联盟hid APIRequest
alitrip.hotel.alliance.hid.get

获取符合条件的菲住联盟hid，目前支持指定日期上线的菲住联盟hid查询
*/
type AlitripHotelAllianceHidGetRequest struct {
    model.Params

    // 查询入参
    allianceInfoRequest   *AllianceInfoRequest 

}

func NewAlitripHotelAllianceHidGetRequest() *AlitripHotelAllianceHidGetRequest{
    return &AlitripHotelAllianceHidGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripHotelAllianceHidGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.alliance.hid.get"
}

func (r AlitripHotelAllianceHidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripHotelAllianceHidGetRequest) SetAllianceInfoRequest(allianceInfoRequest *AllianceInfoRequest) error {
    r.allianceInfoRequest = allianceInfoRequest
    r.Set("alliance_info_request", allianceInfoRequest)
    return nil
}

func (r AlitripHotelAllianceHidGetRequest) GetAllianceInfoRequest() *AllianceInfoRequest {
    return r.allianceInfoRequest
}

