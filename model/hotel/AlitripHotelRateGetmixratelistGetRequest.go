package hotel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店评论接口 APIRequest
alitrip.hotel.rate.getmixratelist.get

酒店评论接口
*/
type AlitripHotelRateGetmixratelistGetRequest struct {
    model.Params

    // 评论参数
    paramGetMixRateListParam   *GetMixRateListParam 

}

func NewAlitripHotelRateGetmixratelistGetRequest() *AlitripHotelRateGetmixratelistGetRequest{
    return &AlitripHotelRateGetmixratelistGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripHotelRateGetmixratelistGetRequest) GetApiMethodName() string {
    return "alitrip.hotel.rate.getmixratelist.get"
}

func (r AlitripHotelRateGetmixratelistGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripHotelRateGetmixratelistGetRequest) SetParamGetMixRateListParam(paramGetMixRateListParam *GetMixRateListParam) error {
    r.paramGetMixRateListParam = paramGetMixRateListParam
    r.Set("param_get_mix_rate_list_param", paramGetMixRateListParam)
    return nil
}

func (r AlitripHotelRateGetmixratelistGetRequest) GetParamGetMixRateListParam() *GetMixRateListParam {
    return r.paramGetMixRateListParam
}

