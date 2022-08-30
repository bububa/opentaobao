package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest 案场活动楼盘维护 API请求
// alibaba.alihouse.newhome.casefield.activity.project.submit
//
// 案场活动楼盘维护
type AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest struct {
	model.Params
	// 请求对象
	_caseFieldActivityProjectsDto *CaseFieldActivityProjectsDto
}

// NewAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitRequest 初始化AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeCasefieldActivityProjectSubmitRequest() *AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest {
	return &AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.casefield.activity.project.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCaseFieldActivityProjectsDto is CaseFieldActivityProjectsDto Setter
// 请求对象
func (r *AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest) SetCaseFieldActivityProjectsDto(_caseFieldActivityProjectsDto *CaseFieldActivityProjectsDto) error {
	r._caseFieldActivityProjectsDto = _caseFieldActivityProjectsDto
	r.Set("case_field_activity_projects_dto", _caseFieldActivityProjectsDto)
	return nil
}

// GetCaseFieldActivityProjectsDto CaseFieldActivityProjectsDto Getter
func (r AlibabaAlihouseNewhomeCasefieldActivityProjectSubmitAPIRequest) GetCaseFieldActivityProjectsDto() *CaseFieldActivityProjectsDto {
	return r._caseFieldActivityProjectsDto
}
