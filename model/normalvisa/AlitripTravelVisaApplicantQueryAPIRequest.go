package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelVisaApplicantQueryAPIRequest 签证申请人查询接口 API请求
// alitrip.travel.visa.applicant.query
//
// 签证申请人查询接口，商家可根据条件查询申请人id，用于签证办理
type AlitripTravelVisaApplicantQueryAPIRequest struct {
	model.Params
	// 请求参数
	_param0 *QueryApplicantParam
}

// NewAlitripTravelVisaApplicantQueryRequest 初始化AlitripTravelVisaApplicantQueryAPIRequest对象
func NewAlitripTravelVisaApplicantQueryRequest() *AlitripTravelVisaApplicantQueryAPIRequest {
	return &AlitripTravelVisaApplicantQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelVisaApplicantQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.visa.applicant.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelVisaApplicantQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTravelVisaApplicantQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求参数
func (r *AlitripTravelVisaApplicantQueryAPIRequest) SetParam0(_param0 *QueryApplicantParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlitripTravelVisaApplicantQueryAPIRequest) GetParam0() *QueryApplicantParam {
	return r._param0
}
