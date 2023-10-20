package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIRequest 天猫服务平台商家投诉单问题列表查询 API请求
// tmall.servicecenter.anomalyrecourse.homedecoration.questioncode.query
//
// 天猫服务平台商家投诉单问题列表查询
type TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIRequest struct {
	model.Params
	// 服务code
	_serviceCode string
}

// NewTmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryRequest 初始化TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIRequest对象
func NewTmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryRequest() *TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIRequest {
	return &TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.anomalyrecourse.homedecoration.questioncode.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceCode is ServiceCode Setter
// 服务code
func (r *TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// GetServiceCode ServiceCode Getter
func (r TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIRequest) GetServiceCode() string {
	return r._serviceCode
}
