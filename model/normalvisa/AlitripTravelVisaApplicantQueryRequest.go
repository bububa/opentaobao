package normalvisa

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
签证申请人查询接口 API请求
alitrip.travel.visa.applicant.query

签证申请人查询接口，商家可根据条件查询申请人id，用于签证办理
*/
type AlitripTravelVisaApplicantQueryAPIRequest struct {
    model.Params
    // 请求参数
    _param0   *QueryApplicantParam
}

// 初始化AlitripTravelVisaApplicantQueryAPIRequest对象
func NewAlitripTravelVisaApplicantQueryRequest() *AlitripTravelVisaApplicantQueryAPIRequest{
    return &AlitripTravelVisaApplicantQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTravelVisaApplicantQueryAPIRequest) GetApiMethodName() string {
    return "alitrip.travel.visa.applicant.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTravelVisaApplicantQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 请求参数
func (r *AlitripTravelVisaApplicantQueryAPIRequest) SetParam0(_param0 *QueryApplicantParam) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlitripTravelVisaApplicantQueryAPIRequest) GetParam0() *QueryApplicantParam {
    return r._param0
}
