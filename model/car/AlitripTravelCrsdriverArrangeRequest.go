package car

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
CRS接送机商家派司机接口 APIRequest
alitrip.travel.crsdriver.arrange

提供给CRS接送机商家派司机的API
*/
type AlitripTravelCrsdriverArrangeRequest struct {
    model.Params

    // 请求对象
    crsDriverArrangeParam   *CrsDriverArrangeParam 

}

func NewAlitripTravelCrsdriverArrangeRequest() *AlitripTravelCrsdriverArrangeRequest{
    return &AlitripTravelCrsdriverArrangeRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTravelCrsdriverArrangeRequest) GetApiMethodName() string {
    return "alitrip.travel.crsdriver.arrange"
}

func (r AlitripTravelCrsdriverArrangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTravelCrsdriverArrangeRequest) SetCrsDriverArrangeParam(crsDriverArrangeParam *CrsDriverArrangeParam) error {
    r.crsDriverArrangeParam = crsDriverArrangeParam
    r.Set("crs_driver_arrange_param", crsDriverArrangeParam)
    return nil
}

func (r AlitripTravelCrsdriverArrangeRequest) GetCrsDriverArrangeParam() *CrsDriverArrangeParam {
    return r.crsDriverArrangeParam
}

