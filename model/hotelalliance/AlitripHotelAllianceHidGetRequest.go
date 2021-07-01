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
type AlitripHotelAllianceHidGetAPIRequest struct {
    model.Params
    // 查询入参
    _allianceInfoRequest   *AllianceInfoRequest
}

// 初始化AlitripHotelAllianceHidGetAPIRequest对象
func NewAlitripHotelAllianceHidGetRequest() *AlitripHotelAllianceHidGetAPIRequest{
    return &AlitripHotelAllianceHidGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelAllianceHidGetAPIRequest) GetApiMethodName() string {
    return "alitrip.hotel.alliance.hid.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelAllianceHidGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AllianceInfoRequest Setter
// 查询入参
func (r *AlitripHotelAllianceHidGetAPIRequest) SetAllianceInfoRequest(_allianceInfoRequest *AllianceInfoRequest) error {
    r._allianceInfoRequest = _allianceInfoRequest
    r.Set("alliance_info_request", _allianceInfoRequest)
    return nil
}

// AllianceInfoRequest Getter
func (r AlitripHotelAllianceHidGetAPIRequest) GetAllianceInfoRequest() *AllianceInfoRequest {
    return r._allianceInfoRequest
}
