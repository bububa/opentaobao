package hotelalliance

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取联盟hid API请求
alitrip.hotel.alliance.hid.get

获取符合条件的菲住联盟hid，目前支持指定日期上线的菲住联盟hid查询
*/
type AlitripHotelAllianceHidGetRequest struct {
    model.Params
    // 查询入参
    allianceInfoRequest   *AllianceInfoRequest
}

// 初始化AlitripHotelAllianceHidGetRequest对象
func NewAlitripHotelAllianceHidGetRequest() *AlitripHotelAllianceHidGetRequest{
    return &AlitripHotelAllianceHidGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelAllianceHidGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.alliance.hid.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelAllianceHidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AllianceInfoRequest Setter
// 查询入参
func (r *AlitripHotelAllianceHidGetRequest) SetAllianceInfoRequest(allianceInfoRequest *AllianceInfoRequest) error {
    r.allianceInfoRequest = allianceInfoRequest
    r.Set("alliance_info_request", allianceInfoRequest)
    return nil
}

// AllianceInfoRequest Getter
func (r AlitripHotelAllianceHidGetRequest) GetAllianceInfoRequest() *AllianceInfoRequest {
    return r.allianceInfoRequest
}
