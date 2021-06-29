package hotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店评论接口 API请求
alitrip.hotel.rate.getmixratelist.get

酒店评论接口
*/
type AlitripHotelRateGetmixratelistGetRequest struct {
    model.Params
    // 评论参数
    _paramGetMixRateListParam   *GetMixRateListParam
}

// 初始化AlitripHotelRateGetmixratelistGetRequest对象
func NewAlitripHotelRateGetmixratelistGetRequest() *AlitripHotelRateGetmixratelistGetRequest{
    return &AlitripHotelRateGetmixratelistGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripHotelRateGetmixratelistGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.rate.getmixratelist.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripHotelRateGetmixratelistGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamGetMixRateListParam Setter
// 评论参数
func (r *AlitripHotelRateGetmixratelistGetRequest) SetParamGetMixRateListParam(_paramGetMixRateListParam *GetMixRateListParam) error {
    r._paramGetMixRateListParam = _paramGetMixRateListParam
    r.Set("param_get_mix_rate_list_param", _paramGetMixRateListParam)
    return nil
}

// ParamGetMixRateListParam Getter
func (r AlitripHotelRateGetmixratelistGetRequest) GetParamGetMixRateListParam() *GetMixRateListParam {
    return r._paramGetMixRateListParam
}
