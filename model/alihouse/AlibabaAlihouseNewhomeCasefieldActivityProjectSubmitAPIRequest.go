package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomecasefieldactivityprojectsubmitAPIRequest 案场活动楼盘维护 API请求
// alibaba.alihouse.newhome.casefield.activity.project.submit
//
// 案场活动楼盘维护
type AlibabaalihousenewhomecasefieldactivityprojectsubmitAPIRequest struct {
	model.Params
	// 请求对象
	_caseFieldActivityProjectsDto *CaseFieldActivityProjectsDto
}

// NewAlibabaalihousenewhomecasefieldactivityprojectsubmitRequest 初始化AlibabaalihousenewhomecasefieldactivityprojectsubmitAPIRequest对象
func NewAlibabaalihousenewhomecasefieldactivityprojectsubmitRequest() *AlibabaalihousenewhomecasefieldactivityprojectsubmitAPIRequest {
	return &AlibabaalihousenewhomecasefieldactivityprojectsubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomecasefieldactivityprojectsubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.casefield.activity.project.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomecasefieldactivityprojectsubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomecasefieldactivityprojectsubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCaseFieldActivityProjectsDto is CaseFieldActivityProjectsDto Setter
// 请求对象
func (r *AlibabaalihousenewhomecasefieldactivityprojectsubmitAPIRequest) SetCaseFieldActivityProjectsDto(_caseFieldActivityProjectsDto *CaseFieldActivityProjectsDto) error {
	r._caseFieldActivityProjectsDto = _caseFieldActivityProjectsDto
	r.Set("case_field_activity_projects_dto", _caseFieldActivityProjectsDto)
	return nil
}

// GetCaseFieldActivityProjectsDto CaseFieldActivityProjectsDto Getter
func (r AlibabaalihousenewhomecasefieldactivityprojectsubmitAPIRequest) GetCaseFieldActivityProjectsDto() *CaseFieldActivityProjectsDto {
	return r._caseFieldActivityProjectsDto
}
