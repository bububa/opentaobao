package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptravelvisaapplicantqueryAPIRequest 签证申请人查询接口 API请求
// alitrip.travel.visa.applicant.query
//
// 签证申请人查询接口，商家可根据条件查询申请人id，用于签证办理
type AlitriptravelvisaapplicantqueryAPIRequest struct {
	model.Params
	// 请求参数
	_param0 *QueryApplicantParam
}

// NewAlitriptravelvisaapplicantqueryRequest 初始化AlitriptravelvisaapplicantqueryAPIRequest对象
func NewAlitriptravelvisaapplicantqueryRequest() *AlitriptravelvisaapplicantqueryAPIRequest {
	return &AlitriptravelvisaapplicantqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptravelvisaapplicantqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.visa.applicant.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptravelvisaapplicantqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptravelvisaapplicantqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求参数
func (r *AlitriptravelvisaapplicantqueryAPIRequest) SetParam0(_param0 *QueryApplicantParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlitriptravelvisaapplicantqueryAPIRequest) GetParam0() *QueryApplicantParam {
	return r._param0
}
