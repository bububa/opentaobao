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
type AlitripHotelHstdfShotelMatchshotelselfAPIRequest struct {
    model.Params
    // HotelMatchParam
    _param0   *HotelMatchParam
}

// 初始化AlitripHotelHstdfShotelMatchshotelselfAPIRequest对象
func NewAlitripHotelHstdfShotelMatchshotelselfRequest() *AlitripHotelHstdfShotelMatchshotelselfAPIRequest{
    return &AlitripHotelHstdfShotelMatchshotelselfAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfShotelMatchshotelselfAPIRequest) GetApiMethodName() string {
    return "alitrip.hotel.hstdf.shotel.matchshotelself"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfShotelMatchshotelselfAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// HotelMatchParam
func (r *AlitripHotelHstdfShotelMatchshotelselfAPIRequest) SetParam0(_param0 *HotelMatchParam) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlitripHotelHstdfShotelMatchshotelselfAPIRequest) GetParam0() *HotelMatchParam {
    return r._param0
}
