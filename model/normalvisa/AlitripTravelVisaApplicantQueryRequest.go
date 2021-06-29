package normalvisa

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
签证申请人查询接口 APIRequest
alitrip.travel.visa.applicant.query

签证申请人查询接口，商家可根据条件查询申请人id，用于签证办理
*/
type AlitripTravelVisaApplicantQueryRequest struct {
    model.Params

    // 请求参数
    param0   *QueryApplicantParam 

}

func NewAlitripTravelVisaApplicantQueryRequest() *AlitripTravelVisaApplicantQueryRequest{
    return &AlitripTravelVisaApplicantQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTravelVisaApplicantQueryRequest) GetApiMethodName() string {
    return "alitrip.travel.visa.applicant.query"
}

func (r AlitripTravelVisaApplicantQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTravelVisaApplicantQueryRequest) SetParam0(param0 *QueryApplicantParam) error {
    r.param0 = param0
    r.Set("param0", param0)
    return nil
}

func (r AlitripTravelVisaApplicantQueryRequest) GetParam0() *QueryApplicantParam {
    return r.param0
}

