package hotelhstdf

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自主匹配标准酒店以及卖家酒店 API请求
alitrip.hotel.hstdf.shotel.matchshotelself

商家通过指定的标准酒店id和卖家酒店id进行匹配
*/
type AlitripHotelHstdfShotelMatchshotelselfRequest struct {
    model.Params
    // HotelMatchParam
    param0   *HotelMatchParam
}

// 初始化AlitripHotelHstdfShotelMatchshotelselfRequest对象
func NewAlitripHotelHstdfShotelMatchshotelselfRequest() *AlitripHotelHstdfShotelMatchshotelselfRequest{
    return &AlitripHotelHstdfShotelMatchshotelselfRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfShotelMatchshotelselfRequest) GetApiMethodName() string {
    return "alitrip.hotel.hstdf.shotel.matchshotelself"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfShotelMatchshotelselfRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// HotelMatchParam
func (r *AlitripHotelHstdfShotelMatchshotelselfRequest) SetParam0(param0 *HotelMatchParam) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

// Param0 Getter
func (r AlitripHotelHstdfShotelMatchshotelselfRequest) GetParam0() *HotelMatchParam {
    return r.param0
}
