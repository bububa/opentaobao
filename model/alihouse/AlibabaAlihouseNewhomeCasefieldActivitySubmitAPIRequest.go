package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomecasefieldactivitysubmitAPIRequest 案场活动维护 API请求
// alibaba.alihouse.newhome.casefield.activity.submit
//
// 案场活动维护
type AlibabaalihousenewhomecasefieldactivitysubmitAPIRequest struct {
	model.Params
	// 请求对象
	_caseFieldActivityDto *CaseFieldActivityDto
}

// NewAlibabaalihousenewhomecasefieldactivitysubmitRequest 初始化AlibabaalihousenewhomecasefieldactivitysubmitAPIRequest对象
func NewAlibabaalihousenewhomecasefieldactivitysubmitRequest() *AlibabaalihousenewhomecasefieldactivitysubmitAPIRequest {
	return &AlibabaalihousenewhomecasefieldactivitysubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomecasefieldactivitysubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.casefield.activity.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomecasefieldactivitysubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomecasefieldactivitysubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCaseFieldActivityDto is CaseFieldActivityDto Setter
// 请求对象
func (r *AlibabaalihousenewhomecasefieldactivitysubmitAPIRequest) SetCaseFieldActivityDto(_caseFieldActivityDto *CaseFieldActivityDto) error {
	r._caseFieldActivityDto = _caseFieldActivityDto
	r.Set("case_field_activity_dto", _caseFieldActivityDto)
	return nil
}

// GetCaseFieldActivityDto CaseFieldActivityDto Getter
func (r AlibabaalihousenewhomecasefieldactivitysubmitAPIRequest) GetCaseFieldActivityDto() *CaseFieldActivityDto {
	return r._caseFieldActivityDto
}
