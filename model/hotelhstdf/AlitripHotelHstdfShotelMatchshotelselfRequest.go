package hotelhstdf

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自主匹配标准酒店以及卖家酒店 APIRequest
alitrip.hotel.hstdf.shotel.matchshotelself

商家通过指定的标准酒店id和卖家酒店id进行匹配
*/
type AlitripHotelHstdfShotelMatchshotelselfRequest struct {
    model.Params

    // HotelMatchParam
    param0   *HotelMatchParam 

}

func NewAlitripHotelHstdfShotelMatchshotelselfRequest() *AlitripHotelHstdfShotelMatchshotelselfRequest{
    return &AlitripHotelHstdfShotelMatchshotelselfRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripHotelHstdfShotelMatchshotelselfRequest) GetApiMethodName() string {
    return "alitrip.hotel.hstdf.shotel.matchshotelself"
}

func (r AlitripHotelHstdfShotelMatchshotelselfRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripHotelHstdfShotelMatchshotelselfRequest) SetParam0(param0 *HotelMatchParam) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlitripHotelHstdfShotelMatchshotelselfRequest) GetParam0() *HotelMatchParam {
    return r.param0
}

